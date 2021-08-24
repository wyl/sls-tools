package core

import (
	"fmt"
	"github.com/BurntSushi/toml"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/mitchellh/go-homedir"
	"os"
	"sls-tools/logger"
)

type Project struct {
	ProjectName string
	LogStore    string
	Index       sls.Index
}

func NewProject(p string) *Project {
	result := Project{}

	path, _ := homedir.Expand(p)
	if _, err := toml.DecodeFile(path, &result); err != nil {
		panic(err)
	}
	return &result
}

type IndexClient struct {
	DumpProject  string
	DumpLogstore []string

	ReceiveProject  string
	ReceiveLogstore [] string

	SlsClient sls.ClientInterface
}

func NewIndexClient() *IndexClient {
	ic := IndexClient{}

	if _, err := toml.DecodeFile("config.toml", &ic); err != nil {
		panic(err)
	}

	auth := NewAliSecret()

	ic.SlsClient = sls.CreateNormalInterface(auth.Endpoint, auth.AccessKeyID, auth.AccessKeySecret, "")
	return &ic
}

func getConfigurationPath(logstore string) string {
	return fmt.Sprintf("./logstore/%v.toml", logstore)
}

func (ic *IndexClient) DumpConfiguration(projectName, logstoreName string) {
	configurationPath := getConfigurationPath(logstoreName)
	if project, err := ic.SlsClient.GetProject(projectName); err != nil {
		logger.MyLogger.ErrorLog.Fatalln(err)
	} else {
		if logstore, err := project.GetLogStore(logstoreName); err != nil {
			logger.MyLogger.ErrorLog.Fatalln(err)
		} else {

			if index, err := logstore.GetIndex(); err != nil {
				logger.MyLogger.ErrorLog.Fatalln(err)
			} else {
				if f, err := os.Create(configurationPath); err != nil {
					defer f.Close()
					logger.MyLogger.ErrorLog.Fatalln(err)
				} else {
					if err := toml.NewEncoder(f).Encode(Project{
						ProjectName: projectName,
						LogStore:    logstoreName,
						Index:       *index,
					}); err != nil {
						logger.MyLogger.ErrorLog.Fatalln(err)
					}
				}
			}

		}
	}

	logger.MyLogger.InfoLog.Println(fmt.Sprintf("%v -> %v 备份完毕。", projectName, logstoreName))
}

func (ic *IndexClient) ApplyConfiguration(projectName, logstoreName string) {
	configurationPath := getConfigurationPath(logstoreName)

	proj := NewProject(configurationPath)
	proj.ProjectName = projectName

	logger.MyLogger.InfoLog.Println(fmt.Sprintf("正在执行 %v -> %v。", proj.ProjectName, proj.LogStore))

	if project, err := ic.SlsClient.GetProject(proj.ProjectName); err != nil {
		logger.MyLogger.ErrorLog.Fatalln(err)
	} else {
		if exit, err := project.CheckLogstoreExist(proj.LogStore); err != nil {
			logger.MyLogger.ErrorLog.Fatalln(err)
		} else {
			if !exit {
				if logger.YesNo(fmt.Sprintf("【%v】 不存在 logstore 【%v】，是否创建 logstore ？", proj.ProjectName, proj.LogStore)) {
					if err := project.CreateLogStore(proj.LogStore, 3650, 2, true, 64); err != nil {
						logger.MyLogger.ErrorLog.Fatalln(err)
					}
				}
			}
		}

		logstore, err := project.GetLogStore(proj.LogStore)
		if err != nil {
			logger.MyLogger.ErrorLog.Fatalln(err)
		}
		if exit, err := logstore.CheckIndexExist(); err != nil {
			logger.MyLogger.ErrorLog.Fatalln(err)
		} else {
			if exit {

				if logger.YesNo(fmt.Sprintf(fmt.Sprintf("%v -> %v 已存在索引，YES 追加索引，NO 覆盖索引。", proj.ProjectName, proj.LogStore))) {
					if originIndex, err := logstore.GetIndex(); err != nil {
						logger.MyLogger.ErrorLog.Fatalln(err)
					} else {
						for key := range proj.Index.Keys {
							_, o := originIndex.Keys[key]
							if !o {
								originIndex.Keys[key] = proj.Index.Keys[key]
							}
						}
						if err := logstore.UpdateIndex(*originIndex); err != nil {
							logger.MyLogger.ErrorLog.Fatalln(err)
						}
					}
					logger.MyLogger.InfoLog.Println(fmt.Sprintf("%v -> %v 追加索引！", proj.ProjectName, proj.LogStore))
				} else {
					if err := logstore.UpdateIndex(proj.Index); err != nil {
						logger.MyLogger.ErrorLog.Fatalln(err)
					}
					logger.MyLogger.InfoLog.Println(fmt.Sprintf("%v -> %v 覆盖索引！", proj.ProjectName, proj.LogStore))
				}
			} else {
				if err := logstore.CreateIndex(proj.Index); err != nil {
					logger.MyLogger.ErrorLog.Fatalln(err)
				}
				logger.MyLogger.InfoLog.Println(fmt.Sprintf("%v -> %v 创建索引！", proj.ProjectName, proj.LogStore))
			}
		}
	}
}

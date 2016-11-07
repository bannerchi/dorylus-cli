package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
	"errors"
	"github.com/bannerchi/dorylus-cli/client"
	"encoding/json"
)

func main(){
	var host string
	var port string
	cli.VersionFlag = cli.BoolFlag{
		Name: "version, V",
		Usage: "print only the version",
	}

	app := cli.NewApp()
	app.Name = "dorylus"
	app.Usage = "A cli for dorylus"
	app.UsageText = "dorylus-cli command tools for dorylus"
	app.HelpName = "a cli for dorylus"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ron Chi",
			Email: "bannerchi@gmail.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "start",
			Aliases: []string{"s"},
			Usage: "start a cron",
			Action: func(c *cli.Context) error {
				var taskString string
				var taskId int
				if c.NArg() > 0 {
					taskString = c.Args().Get(0)
					taskId, _ = strconv.Atoi(taskString)

					if host != "" {
						res, err := client.RunJobByIdAndServerId(taskId, host + ":" + port)
						if err != nil {
							fmt.Printf("error occur %s \n", err)
						} else {
							fmt.Printf("start task %d %s\n", taskId, res)
						}

					} else {
						return errors.New("no host")
					}

				} else {
					return errors.New("error args")
				}

				return nil
			},
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "host",
					Value: "127.0.0.1",
					Usage: "dorylus server host",
					Destination: &host,
				},
				cli.StringFlag{
					Name: "port",
					Value: "8989",
					Usage: "dorylus server port",
					Destination: &port,
				},
			},
		},
		{
			Name: "remove",
			Aliases: []string{"r"},
			Usage: "remove a cron",
			Action: func(c *cli.Context) error {
				var taskString string
				var taskId int
				if c.NArg() > 0 {
					taskString = c.Args().Get(0)
					taskId, _ = strconv.Atoi(taskString)

					if host != "" {
						res, err := client.RmJobByIdAndServerId(taskId, host + ":" + port)
						if err != nil {
							return err
						} else {
							fmt.Printf("remove task %d %s\n", taskId, res)
						}

					} else {
						return errors.New("no host")
					}

				} else {
					return errors.New("error args")
				}

				return nil
			},
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "host",
					Value: "127.0.0.1",
					Usage: "dorylus server host",
					Destination: &host,
				},
				cli.StringFlag{
					Name: "port",
					Value: "8989",
					Usage: "dorylus server port",
					Destination: &port,
				},
			},

		},
		{
			Name: "rtr",
			Aliases: []string{"t"},
			Usage: "get ready to run cron job",
			Action: func(c *cli.Context) error {
				var size int
				if c.NArg() > 0 {
					sizeString := c.Args().Get(0)
					size, _ = strconv.Atoi(sizeString)

					if host != "" {
						entries, err := client.GetReadyToRunJob(size, host + ":" + port);
						if err != nil {
							return err
						} else {
							if len(entries) > 0 {
								for _, e := range entries {
									var processStatus string
									var states client.ProcessState
									fmt.Println("--------------\n")
									fmt.Printf(" task_id: %d \n pid: %d \n name: %s \n status: %d \n", e.Tid, e.Pid, e.Name, e.Status)
									processStatus, _ = client.GetProcStatusByPid(e.Pid, host + ":" + port)
									json.Unmarshal([]byte(processStatus), &states)
									fmt.Printf("****pid %d****\n", states.Pid)
									fmt.Printf(" pid: %d \n isRunning: %v \n cpuPercent: %f \n memPercent: %f\n",
										states.Pid,states.IsRunning, states.CpuPercent, states.MemoryPercent)
									fmt.Println("**************\n")
									fmt.Println("--------------\n")
								}
							}
						}

					} else {
						return errors.New("no host")
					}

				} else {
					return errors.New("error args")
				}

				return nil
			},
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "host",
					Value: "127.0.0.1",
					Usage: "dorylus server host",
					Destination: &host,
				},
				cli.StringFlag{
					Name: "port",
					Value: "8989",
					Usage: "dorylus server port",
					Destination: &port,
				},
			},

		},
	}

	app.Run(os.Args)
}


package main

import "fmt"

type choice struct {
	cmd         string
	description string
	strorynode  *storynode
}

type storynode struct {
	text    string
	choices []*choice
}

func main() {
	start := storynode{
		text: "you are in large chamber, deep underground you see three passages leading out",
		choices: []*choice{
			{
				"N",
				"Darkness",
				&storynode{
					text: "It is dark,pitch black you can't see a thing",
					choices: []*choice{
						&choice{
							"A",
							"Lantern ",
							&storynode{
								text: "The dark passage is lit by your lantern.",
								choices: []*choice{
									&choice{
										cmd:         "A",
										description: "Go North",
										strorynode: &storynode{
											text: "big fat bear sitting in front of the next door.you have to fight with him.",
											choices: []*choice{&choice{
												cmd:         "A",
												description: "knif",
												strorynode: &storynode{
													text:    "too small,you cannot fight with it. loose.",
													choices: nil,
												},
											},
												&choice{
													cmd:         "B",
													description: "sword",
													strorynode: &storynode{
														text:    "greate choice. you win. congratulation",
														choices: nil,
													},
												},
												&choice{
													cmd:         "C",
													description: "gun with only one bullet",
													strorynode: &storynode{
														text:    "poor choice. you cannot kill bear with one bullet. you loose",
														choices: nil,
													},
												},
											},
										},
									},
									&choice{
										cmd:         "B",
										description: "Go South",
										strorynode: &storynode{
											text: "In the south, long rever with dangerous animal. back water. you have to select one option so you can servive. and cross the rever. ",
											choices: []*choice{
												&choice{
													cmd:         "A",
													description: "boat",
													strorynode: &storynode{
														text:    "while stumbling around in the revere,you are eaten by shark",
														choices: nil,
													},
												},
												&choice{
													cmd:         "B",
													description: "submarine",
													strorynode: &storynode{
														text: "great choice. Please sit down....ohh No..while going , big rock come into your ways." +
															" you have to break that rock. otherwise it will destroy your submarine",
														choices: []*choice{
															&choice{
																cmd:         "A",
																description: "bomb",
																strorynode: &storynode{
																	text:    "you win.",
																	choices: nil,
																},
															},
															&choice{
																cmd:         "B",
																description: "missile",
																strorynode: &storynode{
																	text:    "you win.",
																	choices: nil,
																},
															},
														},
													},
												},
												&choice{
													cmd:         "c",
													description: "lifejacket",
													strorynode: &storynode{
														text:    "while stumbling around in the revere,you are eaten by crocodile",
														choices: nil,
													},
												},
											},
										},
									},
								},
							},
						},
						&choice{
							cmd:         "B",
							description: "Go back",
							strorynode:  nil,
						},
					},
				},
			},
			{
				"S",
				"next room",
				&storynode{
					text:    "",
					choices: nil,
				},
			},
			{
				"E",
				"trap",
				&storynode{
					text:    "",
					choices: nil,
				},
			},
		},
	}
	fmt.Println(start)

}

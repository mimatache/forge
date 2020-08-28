package manifest

func InitializeDemoForge() Forge {
	return Forge{
		Include: []string{},
		Forgeries: Forgeries{
			{
				Name:        "shouldSucceed",
				Description: "A forgery that should succeed",
				Pre:         nil,
				Cmd:         "echo Success",
			},
			{
				Name:        "anotherSuccess",
				Description: "Another forgery that should succeed",
				Pre:         nil,
				Cmd:         "forge",
			},
			{
				Name:        "shouldFail",
				Description: "A forgery that should always fail",
				Pre:         nil,
				Cmd:         "exit 1",
			},
			{
				Name:        "chainForge",
				Description: "A forgery that is composed only out of other forge executions",
				Pre:         []ForgeryName{"shouldSucceed", "anotherSuccess"},
			},
		},
	}
}

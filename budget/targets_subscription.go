package main

func subscriptionTarget() []Targets {
	var netflix Targets
	netflix.Name = "netflix"
	netflix.Cat = []string{con, enter, sub}
	netflix.Variant = []string{"netflix"}

	var audible Targets
	audible.Name = "audible"
	audible.Cat = []string{con, enter, sub}
	audible.Variant = []string{"audible"}

	var apple Targets
	apple.Name = "apple tv"
	apple.Cat = []string{con, enter, sub}
	apple.Variant = []string{"apple"}

	var github Targets
	github.Name = "github"
	github.Cat = []string{con, sub}
	github.Variant = []string{"github"}

	var microsoft Targets
	microsoft.Name = "microsoft"
	microsoft.Cat = []string{con, sub}
	microsoft.Variant = []string{"microsoft"}

	var scmp Targets
	scmp.Name = "scmp"
	scmp.Cat = []string{sub, con}
	scmp.Variant = []string{"scmp hong kong"}

	var wsj Targets
	wsj.Name = "wsj"
	wsj.Cat = []string{sub, con}
	wsj.Variant = []string{"wall-st-journa"}

	var youtube Targets
	youtube.Name = "youtube"
	youtube.Cat = []string{sub, con}
	youtube.Variant = []string{"youtubepre"}

	var sub = []Targets{}
	sub = append(sub, netflix, audible, apple, github, scmp, wsj, youtube)
	return sub
}

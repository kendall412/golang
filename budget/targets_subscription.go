package main

func subscriptionTarget() []Targets {
	var netflix Targets
	netflix.Name = "netflix"
	netflix.Cat = []string{cats["con"], cats["enter"], cats["sub"]}
	netflix.Variant = []string{"netflix"}

	var audible Targets
	audible.Name = "audible"
	audible.Cat = []string{cats["con"], cats["enter"], cats["sub"]}
	audible.Variant = []string{"audible"}

	var apple Targets
	apple.Name = "apple tv"
	apple.Cat = []string{cats["con"], cats["enter"], cats["sub"]}
	apple.Variant = []string{"apple"}

	var github Targets
	github.Name = "github"
	github.Cat = []string{cats["con"], cats["sub"]}
	github.Variant = []string{"github"}

	var microsoft Targets
	microsoft.Name = "microsoft"
	microsoft.Cat = []string{cats["con"], cats["sub"]}
	microsoft.Variant = []string{"microsoft"}

	var scmp Targets
	scmp.Name = "scmp"
	scmp.Cat = []string{cats["con"], cats["sub"]}
	scmp.Variant = []string{"scmp hong kong"}

	var wsj Targets
	wsj.Name = "wsj"
	wsj.Cat = []string{cats["con"], cats["sub"]}
	wsj.Variant = []string{"wall-st-journa"}

	var youtube Targets
	youtube.Name = "youtube"
	youtube.Cat = []string{cats["con"], cats["sub"]}
	youtube.Variant = []string{"youtubepre"}

	var sub = []Targets{}
	sub = append(sub, netflix, audible, apple, github, microsoft, scmp, wsj, youtube)
	return sub
}

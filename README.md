## Get battle results of splatoon2 from "ika-ring"
It is necessary to get the iksm_session key using mitmproxy, for example,
and set it as the environmental variable named IKSM_SESSION.
Only "/api/results/N" and "/api/data/stages" are included at the current stage.
Some unnecessary information is dropped from json obtained by /api/results/N.

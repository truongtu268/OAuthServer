package FakeData

import "github.com/truongtu268/OAuthServer/Model"

var Providers = []Model.Provider{
	Model.Provider{
		Name:     "google",
		Cid:      "251196717827-4oh8dp23ftu0555b905n0coa08lua1km.apps.googleusercontent.com",
		Csecret:  "CWeUqK2JHs_-MefG1A4kG7vt",
		Callback: "http://127.0.0.1:9090/user/auth/google",
		Scope: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Client:   "https://www.googleapis.com/oauth2/v3/userinfo",
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://accounts.google.com/o/oauth2/token",
	},
	Model.Provider{
		Name:     "github",
		Cid:      "7fb34f5e3f8a61d5c148",
		Csecret:  "f769509e38a4c0200b2bf7fe97e98534c3bac581",
		Callback: "http://127.0.0.1:9090/user/auth/github",
		Scope: []string{
			"user",
			"publicrepo",
		},
		Client:   "https://api.github.com/user",
		AuthURL:  "https://github.com/login/oauth/authorize",
		TokenURL: "https://github.com/login/oauth/access_token",
	},
	Model.Provider{
		Name:     "instagram",
		Cid:      "cb1d9f05d80e4a2c821c2916499c4b82",
		Csecret:  "907debf7db834c2594b595de45ae533b",
		Callback: "http://127.0.0.1:9090/user/auth/instagram",
		Scope: []string{
			"basic",
		},
		Client:   "https://api.instagram.com/v1/users/self/media/recent/?access_token=",
		AuthURL:  "https://api.instagram.com/oauth/authorize",
		TokenURL: "https://api.instagram.com/oauth/access_token",
	},
	Model.Provider{
		Name:     "internal",
		Cid:      "cb1d9f05d80e4a2c821c2916499c4b8q",
		Csecret:  "907debf7db834c2594b595de45ae533q",
		Callback: "http://127.0.0.1:9090/user/auth/local_authorize",
		Scope: []string{
			"basic",
			"read profile",
		},
		Client:   "http://127.0.0.1:9090/private/info",
		AuthURL:  "http://127.0.0.1:9090/oauth/authorize",
		TokenURL: "http://127.0.0.1:9090/oauth/access_token",
	},
}

var Clients = []Model.Client{
	Model.Client{
		ID:         "cb1d9f05d80e4a2c821c2916499c4b8q",
		Secret:     "907debf7db834c2594b595de45ae533q",
		Scope:      []string{"basic", "read profile"},
		CallBack:   []string{"http://127.0.0.1:9090/user/auth/local_authorize"},
		TrustLevel: "5",
	},
}

var Users = []Model.User{
	Model.User{
		Name:   "Truongtu",
		Email:  "truogntu268@gmail.com",
		Avatar: "tu dep trai",
		SecurityInfos: []Model.UserSecurityInfo{
			Model.UserSecurityInfo{
				ProviderLogin: "internal",
				Username:      "truongtu",
				Password:      "123456789",
			},
		},
	},
}

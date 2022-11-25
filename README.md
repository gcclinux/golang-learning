# GOLANG-LEARNING
Install Git-bash as terminal for Windows<br>
Install go from https://go.dev/dl/<br>
Install Visual Studio for any platform<br>
After Visual Studio installed need to intall Go modules<br>
	Crtl + Shift + P <br>
	> Go: Install/update (select all)<br>
<br>
Create Project folder<br>
<br>
Enter project folder via Git-Bash command line<br>
Get your GOPATH by following command <br>
	$ echo $GOPTAH<br>
Initiate Go modules in project folder<br>
	$ go mod init 'myapp'  'name of your application or location like github.com/user/myapp'<br>
<br>
Opening Settings short cup<br>
	Crtl + ,<br>
	File --> Preferences --> Settings<br>
	Update Settings<br>
		"terminal.integrated.defaultProfile.windows": "Git Bash",<br>
		"terminal.external.windowsExec": "C:\\Program Files\\Git\\git-bash.exe",<br>
		"terminal.integrated.cwd": "C:\\Users\\ricar\\Desktop\\Nextcloud\\Programming\\GitServer\\golang",<br>
<br>
Download godoc modules for references<br>
	$ go get golang.org/x/tools/cmd/godoc<br>
	$ go install golang.org/x/tools/cmd/godoc<br>
To view godoc initial webserver<br>
	$ godoc --http=:6060<br>
Open favour browser browser and open URL<br>
	http://localhost:6060<br>
Very first line in all GO files start with package<br>
	package main<br>

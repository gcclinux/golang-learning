# GOLANG-LEARNING
Install Git-bash as terminal for Windows<br>
Install go from https://go.dev/dl/<br>
Install Visual Studio for any platform<br>
After Visual Studio installed need to intall Go modules<br>
&nbsp;&nbsp;&nbsp;&nbsp;Crtl + Shift + P <br>
&nbsp;&nbsp;&nbsp;&nbsp;> Go: Install/update (select all)<br>
<br>
Create Project folder<br>
<br>
Enter project folder via Git-Bash command line<br>
Get your GOPATH by following command <br>
&nbsp;&nbsp;&nbsp;&nbsp;$ echo $GOPTAH<br>
Initiate Go modules in project folder<br>
&nbsp;&nbsp;&nbsp;&nbsp;$ go mod init 'myapp'  'name of your application or location like github.com/user/myapp'<br>
<br>
Opening Settings short cup<br>
&nbsp;&nbsp;&nbsp;&nbsp;Crtl + ,<br>
&nbsp;&nbsp;&nbsp;&nbsp;File --> Preferences --> Settings<br>
&nbsp;&nbsp;&nbsp;&nbsp;Update Settings<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"terminal.integrated.defaultProfile.windows": "Git Bash",<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"terminal.external.windowsExec": "C:\\Program Files\\Git\\git-bash.exe",<br>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"terminal.integrated.cwd": "C:\\Users\\ricar\\Desktop\\Nextcloud\\Programming\\GitServer\\golang",<br>
<br>
Download godoc modules for references<br>
&nbsp;&nbsp;&nbsp;&nbsp;$ go get golang.org/x/tools/cmd/godoc<br>
&nbsp;&nbsp;&nbsp;&nbsp;$ go install golang.org/x/tools/cmd/godoc<br>
To view godoc initial webserver<br>
&nbsp;&nbsp;&nbsp;&nbsp;$ godoc --http=:6060<br>
Open favour browser browser and open URL<br>
&nbsp;&nbsp;&nbsp;&nbsp;http://localhost:6060<br>
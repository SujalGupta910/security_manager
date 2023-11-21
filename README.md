# README

## About
SecureU Internship Assignment
## Building
Open the project directory in cmd and run the command `wails build ; .\build\bin\security_manager.exe`. An admin prompt window will pop up, click on yes.
Once the app is open, click on Users > User Icon Image
## wails dev
The application requires admin privileges which is why it can't be run using `wails dev` command.
To do so, find the wails.exe.manifest file at build\windows and comment out the following code like this:
```
    <!-- <trustInfo xmlns="urn:schemas-microsoft-com:asm.v3">
        <security>
            <requestedPrivileges>
                <requestedExecutionLevel level="requireAdministrator" uiAccess="false"/>
            </requestedPrivileges>
        </security>
    </trustInfo> -->
```

# README

## About
SecureU(OPC) Pvt Ltd Internship Assignment
## Running
Locate the security_manager.msi file and download it.
Once downloaded, run it.
Locate the security_manager shortcut on the desktop and double click on it.
Grant the required admin rights, as the app changes registry keys.
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
## Building
`wails build`

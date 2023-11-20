# README

## About
SecureU Internship Assignment
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

To build a redistributable, production mode package, use `wails build`.

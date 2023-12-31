import './style.css';
import './app.css';

import {DisableWindowsUpdates,DisableCmd, DisableDownloads, DisableFacebook, ChangeScreenTimeout} from '../wailsjs/go/main/App';

window.showPage = function (page) {
    document.getElementById('users').style.display = 'none';
    document.getElementById('userDetails').style.display = 'none';
    document.getElementById(page).style.display = 'block';
}

window.showUserDetails = function (userId) {
    document.getElementById('users').style.display = 'none';
    document.getElementById('userDetails').style.display = 'none';
    document.getElementById('userDetails').style.display = 'block';
    document.getElementById('selectedUser').innerText = 'User: ' + userId;

}

window.submitForm = function () {
    const outputContainer = document.getElementById('outputContainer');
    outputContainer.style.display = 'block';
    outputContainer.innerHTML = 'Running.... <br>';

    if (document.getElementById("noAutoUpdates").checked) {
        try{
            DisableWindowsUpdates()
            .then((message,err) => {
                outputContainer.innerHTML += message + '<br>';
                outputContainer.classList.remove('error');
                outputContainer.classList.add('success');
            })
            .catch((err) => {
                console.error('Error disabling Windows Auto Updates', err);
                outputContainer.innerHTML += 'Error disabling Windows Updates'+' '+err+'<br>';
                outputContainer.classList.remove('success');
                outputContainer.classList.add('error');
            });
        } catch (err) {
            console.error('An unexpected error occurred:', err);
            outputContainer.innerHTML += 'An unexpected error occurred:'+' '+err+'<br>';
            outputContainer.classList.remove('success', 'error');
        }
    }
    if (document.getElementById("disableCmd").checked) {
        try{
            DisableCmd()
            .then((mess,err) => {
                outputContainer.innerHTML += mess+ '<br>';
                outputContainer.classList.remove('error');
                outputContainer.classList.add('success');
            })
            .catch((err) => {
                console.error('Error disabling access to CMD:', err);
                outputContainer.innerHTML += 'Error disabling access to CMD'+' '+err+'<br>';
                outputContainer.classList.remove('success');
                outputContainer.classList.add('error');
            });
        } catch (err) {
            console.error('An unexpected error occurred:', err);
            outputContainer.innerHTML += 'An unexpected error occurred:'+' '+err+'<br>';
            outputContainer.classList.remove('success', 'error');
        }
    }
    if (document.getElementById("disableDownloads").checked) {
        try{
            DisableDownloads()
            .then((mess,err) => {
                outputContainer.innerHTML += mess + '<br>';
                outputContainer.classList.remove('error');
                outputContainer.classList.add('success');
            })
            .catch((err) => {
                console.error('Error disabling downloads:', err);
                outputContainer.innerHTML += 'Error disabling downloads'+' '+err+'<br>';
                outputContainer.classList.remove('success');
                outputContainer.classList.add('error');
            });
        } catch (err) {
            console.error('An unexpected error occurred:', err);
            outputContainer.innerHTML += 'An unexpected error occurred:'+' '+err+'<br>';
            outputContainer.classList.remove('success', 'error');
        }
    }
    if (document.getElementById("disableFacebook").checked) {
        try{
            DisableFacebook()
            .then((mess,err) => {
                outputContainer.innerHTML += mess+'<br>';
                outputContainer.classList.remove('error');
                outputContainer.classList.add('success');
            })
            .catch((err) => {
                console.error('Error disabling access to Facebook:', err);
                outputContainer.innerHTML += 'Error disabling access to Facebook'+' '+err+'<br>';
                outputContainer.classList.remove('success');
                outputContainer.classList.add('error');
            });
        } catch (err) {
            console.error('An unexpected error occurred:', err);
            outputContainer.innerHTML += 'An unexpected error occurred:'+' '+err+'<br>';
            outputContainer.classList.remove('success', 'error');
        }
    }
    if (document.getElementById("changeScreenTimeout").checked) {
        try{
            ChangeScreenTimeout()
            .then((mess,err) => {
                outputContainer.innerHTML += mess+ '<br>';
                outputContainer.classList.remove('error');
                outputContainer.classList.add('success');
            })
            .catch((err) => {
                console.error('Error changing timeout:', err);
                outputContainer.innerHTML += 'Error Changing timeout'+' '+err+'<br>';
                outputContainer.classList.remove('success');
                outputContainer.classList.add('error');
            });
        } catch (err) {
            console.error('An unexpected error occurred:', err);
            outputContainer.innerHTML += 'An unexpected error occurred:'+' '+err+'<br>';
            outputContainer.classList.remove('success', 'error');
        }

    }

};

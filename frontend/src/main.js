import './style.css';
import './app.css';

import {DisableWindowsUpdates,DisableCmd, DisableDownloads, DisableFacebook, ChangeScreenTimeout} from '../wailsjs/go/main/App';

// Setup the greet function
window.submitForm = function () {
    if (document.getElementById("noAutoUpdates").checked) {
        try{
            DisableWindowsUpdates()
            .then(() => {
                // Handle success
                const messageElement = document.getElementById('message');
                messageElement.textContent += 'Disabled Windows Auto Updates<br>';
                messageElement.classList.remove('error'); // Remove any previous error class
                messageElement.classList.add('success'); // Add success class
            })
            .catch((err) => {
                // Handle errors
                console.error('Error disabling Windows Auto Updates', err);
                const messageElement = document.getElementById('message');
                messageElement.textContent += 'Error disabling Windows Auto Updates. Please try again.<br>';
                messageElement.classList.remove('success'); // Remove any previous success class
                messageElement.classList.add('error'); // Add error class
            });
        } catch (err) {
            // Handle unexpected errors
            console.error('An unexpected error occurred:', err);
            const messageElement = document.getElementById('message');
            messageElement.textContent = 'An unexpected error occurred. Please try again.';
            messageElement.classList.remove('success', 'error'); // Remove any previous classes
        }
    }
    if (document.getElementById("disableCmd").checked) {
        try{
            DisableCmd()
            .then(() => {
                // Handle success
                const messageElement = document.getElementById('message');
                messageElement.textContent += 'Access to CMD disabled successfully.<br>';
                messageElement.classList.remove('error'); // Remove any previous error class
                messageElement.classList.add('success'); // Add success class
            })
            .catch((err) => {
                // Handle errors
                console.error('Error disabling access to CMD:', err);
                const messageElement = document.getElementById('message');
                messageElement.textContent += 'Error disabling access to CMD. Please try again.<br>';
                messageElement.classList.remove('success'); // Remove any previous success class
                messageElement.classList.add('error'); // Add error class
            });
        } catch (err) {
            // Handle unexpected errors
            console.error('An unexpected error occurred:', err);
            const messageElement = document.getElementById('message');
            messageElement.textContent = 'An unexpected error occurred. Please try again.';
            messageElement.classList.remove('success', 'error'); // Remove any previous classes
        }
    }
    if (document.getElementById("disableDownloads").checked) {
        try{
            DisableDownloads()
            .then(() => {
                // Handle success
                const messageElement = document.getElementById('message');
                messageElement.textContent += 'Downloads in Chrome and Edge disabled successfully.<br>';
                messageElement.classList.remove('error'); // Remove any previous error class
                messageElement.classList.add('success'); // Add success class
            })
            .catch((err) => {
                // Handle errors
                console.error('Error disabling downloads:', err);
                const messageElement = document.getElementById('message');
                messageElement.textContent += 'Error disabling downloads. Please try again.<br>';
                messageElement.classList.remove('success'); // Remove any previous success class
                messageElement.classList.add('error'); // Add error class
            });
        } catch (err) {
            // Handle unexpected errors
            console.error('An unexpected error occurred:', err);
            const messageElement = document.getElementById('message');
            messageElement.textContent = 'An unexpected error occurred. Please try again.';
            messageElement.classList.remove('success', 'error'); // Remove any previous classes
        }
    }
    if (document.getElementById("disableFacebook").checked) {
        try{
            DisableFacebook()
            .then(() => {
                // Handle success
                const messageElement = document.getElementById('message');
                messageElement.textContent += 'Access to Facebook disabled successfully.<br>';
                messageElement.classList.remove('error'); // Remove any previous error class
                messageElement.classList.add('success'); // Add success class
            })
            .catch((err) => {
                // Handle errors
                console.error('Error disabling access to Facebook:', err);
                const messageElement = document.getElementById('message');
                messageElement.textContent += 'Error disabling access to Facebook. Please try again.<br>';
                messageElement.classList.remove('success'); // Remove any previous success class
                messageElement.classList.add('error'); // Add error class
            });
        } catch (err) {
            // Handle unexpected errors
            console.error('An unexpected error occurred:', err);
            const messageElement = document.getElementById('message');
            messageElement.textContent = 'An unexpected error occurred. Please try again.';
            messageElement.classList.remove('success', 'error'); // Remove any previous classes
        }
    }
    if (document.getElementById("changeScreenTimeout").checked) {
        try{
            ChangeScreenTimeout()
            .then(() => {
                // Handle success
                const messageElement = document.getElementById('message');
                messageElement.textContent += 'Changed timeout to 3 minutes successfully.<br>';
                messageElement.classList.remove('error'); // Remove any previous error class
                messageElement.classList.add('success'); // Add success class
            })
            .catch((err) => {
                // Handle errors
                console.error('Error changing timeout:', err);
                const messageElement = document.getElementById('message');
                messageElement.textContent += 'Error Changing timeout. Please try again.<br>';
                messageElement.classList.remove('success'); // Remove any previous success class
                messageElement.classList.add('error'); // Add error class
            });
        } catch (err) {
            // Handle unexpected errors
            console.error('An unexpected error occurred:', err);
            const messageElement = document.getElementById('message');
            messageElement.textContent = 'An unexpected error occurred. Please try again.';
            messageElement.classList.remove('success', 'error'); // Remove any previous classes
        }

    }

};

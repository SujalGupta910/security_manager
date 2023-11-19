import './style.css';
import './app.css';

import {DisableWindowsUpdates,DisableCmd} from '../wailsjs/go/main/App';

// Setup the greet function
window.submitForm = function () {
    if (document.getElementById("noAutoUpdates").checked) {
        try{
            DisableWindowsUpdates()
            .then(() => {
                // Handle success
                const messageElement = document.getElementById('message');
                messageElement.textContent = 'Disabled Windows Auto Updates';
                messageElement.classList.remove('error'); // Remove any previous error class
                messageElement.classList.add('success'); // Add success class
            })
            .catch((err) => {
                // Handle errors
                console.error('Error disabling Windows Auto Updates', err);
                const messageElement = document.getElementById('message');
                messageElement.textContent = 'Error disabling Windows Auto Updates. Please try again.';
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
                messageElement.textContent = 'Access to CMD disabled successfully.';
                messageElement.classList.remove('error'); // Remove any previous error class
                messageElement.classList.add('success'); // Add success class
            })
            .catch((err) => {
                // Handle errors
                console.error('Error disabling access to CMD:', err);
                const messageElement = document.getElementById('message');
                messageElement.textContent = 'Error disabling access to CMD. Please try again.';
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

    }
    if (document.getElementById("disableFacebook").checked) {

    }
    if (document.getElementById("changeScreenTimeout").checked) {

    }

};

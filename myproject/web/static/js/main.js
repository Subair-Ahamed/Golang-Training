/* main.js */

document.addEventListener("DOMContentLoaded", function () {
    // Code to execute when DOM is fully loaded
    console.log("DOM is fully loaded.");

    // Example AJAX request to get user data
    fetch("/user?id=1")
        .then(response => response.json())
        .then(data => {
            console.log("User data:", data);
            displayUserInfo(data);
        })
        .catch(error => console.error("Error fetching user data:", error));

    function displayUserInfo(user) {
        const userInfoContainer = document.querySelector("#user-info");
        userInfoContainer.innerHTML = `
            <div class="user-info">
                <h2>User Information</h2>
                <p><strong>ID:</strong> ${user.id}</p>
                <p><strong>Name:</strong> ${user.name}</p>
            </div>
        `;
    }
});

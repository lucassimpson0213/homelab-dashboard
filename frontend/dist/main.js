"use strict";
// Main wrapper function
async function tryCatch(promise) {
    try {
        const data = await promise;
        return { data, error: null };
    }
    catch (error) {
        return { data: null, error: error };
    }
}
async function handleSubmit(event, form, input) {
    event.preventDefault();
    const data = input?.value ?? "none";
    console.log("This function gets called");
    if (data === "none") {
        console.error("There was no data submitted for the form");
        return;
    }
    if (!form) {
        console.error("No form detected");
        return;
    }
    if (URL.canParse(data)) {
        let fd = new FormData(form);
        let posted = fetch('http://localhost:8080/api/getlinks', {
            method: "POST",
            body: fd
        });
        let fetched = fetch('http://localhost:8080/api/getlinks', {
            method: "POST",
            body: fd
        });
        let result = await tryCatch(fetched);
        if (result.error !== null) {
            let linkList = document.createElement("li");
            linkList.innerText = "There are no elements";
        }
        let currentList = document.querySelector("ul");
        if (currentList == null) {
            console.error("There is no list to remove");
        }
        else {
            currentList.remove();
        }
        let linkList = document.createElement("ul");
        const linkJson = await result.data?.json();
        if (!result.data) {
            linkList.innerText = "default text 404";
        }
        const MINIMUM_URL_LENGTH = 1;
        for (const element of linkJson) {
            console.log(element);
            if (element.url.length > MINIMUM_URL_LENGTH) {
                const linkElement = document.createElement("li");
                linkElement.innerText = element.url;
                linkList.appendChild(linkElement);
            }
        }
        const body = document.querySelector("body");
        body?.appendChild(linkList) ?? console.error("unable to append child to linkList");
        // let value = result.error ? result.error : JSON.parse(result.data.json)
    }
}
document.addEventListener("DOMContentLoaded", () => {
    const form = document.querySelector('#linkform');
    const submitbutton = document.querySelector('#submitbutton');
    const input = document.querySelector('#inputf');
    const body = document.querySelector('body');
    const list = document.createElement("ul");
    form?.addEventListener("submit", async (event) => {
        await handleSubmit(event, form, input);
    });
    body?.appendChild(list) ?? console.error("unable to append unordered list");
    console.log(form);
});
//# sourceMappingURL=main.js.map
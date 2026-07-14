const form = document.querySelector('#linkform');
const submitbutton = document.querySelector('#submitbutton');
const input = document.querySelector('#inputf');
const list = document.createElement("ul");
// Main wrapper function
export async function tryCatch(promise) {
    try {
        const data = await promise;
        return { data, error: null };
    }
    catch (error) {
        return { data: null, error: error };
    }
}
async function handleSubmit(event) {
    event.preventDefault();
    const data = input?.value ?? "none";
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
        let fetched = fetch('http://fedoramac.lan:8080/getlinks');
        let result = await tryCatch(fetched);
        if (result.error !== null) {
            let linkList = document.createElement("li");
            linkList.innerText = "There are no elements";
        }
        // let value = result.error ? result.error : JSON.parse(result.data.json)
    }
}
form?.addEventListener("submit", handleSubmit);
console.log(form);
//# sourceMappingURL=main.js.map
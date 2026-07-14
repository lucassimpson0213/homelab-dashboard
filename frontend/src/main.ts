

const form = document.querySelector<HTMLFormElement>('#linkform')
const submitbutton = document.querySelector('#submitbutton')
const input = document.querySelector<HTMLInputElement>('#inputf')


async function handleSubmit(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    const data = input?.value ?? "none"

    if (data === "none") {
        console.error("There was no data submitted for the form")
        return;
    }

    if (!form) {
        console.error("No form detected");
        return;
    }

    if (URL.canParse(data)) {
        let fd = new FormData(form);

        let fetched = await fetch('http://fedoramac.lan:8080/getlinks');

        console.log(fetched);
    }
}
form?.addEventListener("submit", handleSubmit);
console.log(form)

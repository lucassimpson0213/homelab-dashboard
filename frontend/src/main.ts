

// make your own tree functions and component framework
//
// preorder traversal of tree
type Success<T> = {
    data: T;
    error: null;
};

type Failure<E> = {
    data: null;
    error: E;
};

type Result<T, E = Error> = Success<T> | Failure<E>;

// Main wrapper function
async function tryCatch<T, E = Error>(
    promise: Promise<T>,
): Promise<Result<T, E>> {
    try {
        const data = await promise;
        return { data, error: null };
    } catch (error) {
        return { data: null, error: error as E };
    }
}


async function handleSubmit(event: SubmitEvent, form: HTMLFormElement, input: HTMLInputElement | null): Promise<void> {
    event.preventDefault();
    const data = input?.value ?? "none"
    console.log("This function gets called");
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

        let posted = fetch('http://localhost:8080/api/getlinks', {
            method: "POST",
            body: fd
        });


        let fetched = fetch('http://localhost:8080/api/getlinks', {
            method: "POST",
            body: fd
        });




        let result: Result<Response, Error> = await tryCatch(fetched);

        if (result.error !== null) {
            let linkList = document.createElement("li");
            linkList.innerText = "There are no elements"
        }

        let currentList = document.querySelector("ul");

        if (currentList == null) {
            console.error("There is no list to remove")
        } else {
            currentList.remove();
        }


        let linkList = document.createElement("ul");


        const linkJson = await result.data?.json();
        if (!result.data) {
            linkList.innerText = "default text 404"
        }


        const MINIMUM_URL_LENGTH = 1;
        for (const element of linkJson) {
            console.log(element)
            if (element.url.length > MINIMUM_URL_LENGTH) {

                const linkElement = document.createElement("li");
                linkElement.innerText = element.url;
                linkList.appendChild(linkElement);

            }

        }
        const body = document.querySelector("body");

        body?.appendChild(linkList) ?? console.error("unable to append child to linkList")



        // let value = result.error ? result.error : JSON.parse(result.data.json)
    }
}
document.addEventListener("DOMContentLoaded", () => {
    const form = document.querySelector<HTMLFormElement>('#linkform')
    const submitbutton = document.querySelector('#submitbutton')
    const input = document.querySelector<HTMLInputElement>('#inputf')
    const body = document.querySelector('body');
    const list = document.createElement("ul");
    form?.addEventListener("submit", async (event: SubmitEvent) => {
        await handleSubmit(event, form, input)
    });
    body?.appendChild(list) ?? console.error("unable to append unordered list");
    console.log(form)
});


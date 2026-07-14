

const form = document.querySelector<HTMLFormElement>('#linkform')
const submitbutton = document.querySelector('#submitbutton')
const input = document.querySelector<HTMLInputElement>('#inputf')
const list = document.createElement("ul");

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
export async function tryCatch<T, E = Error>(
    promise: Promise<T>,
): Promise<Result<T, E>> {
    try {
        const data = await promise;
        return { data, error: null };
    } catch (error) {
        return { data: null, error: error as E };
    }
}


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

        let fetched = fetch('http://fedoramac.lan:8080/getlinks');

        let result: Result<Response, Error> = await tryCatch(fetched);
        if (result.error !== null) {
            let linkList = document.createElement("li");
            linkList.innerText = "There are no elements"
        }


        // let value = result.error ? result.error : JSON.parse(result.data.json)
    }
}
form?.addEventListener("submit", handleSubmit);
console.log(form)

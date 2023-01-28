const eventHandler = (element, callback, eventName) => element.addEventListener(eventName, callback);

const click = (e, c) => eventHandler(e, c, "click")
const keyDown = (e, c) => eventHandler(e, c, "input")

const btn = document.getElementById("query");
const results = document.getElementById("results");
const block = document.getElementById("block");
const input = document.getElementById("input");

const textNode = (e) => document.createTextNode(e)
const textTag = (node) => document.createElement("p").appendChild(node)
const showTag = (e) => results.appendChild(textTag(textNode(e)))

const p = () => document.getElementById("warning");  // Allows to check if it is on the page

/* Constitutes an animated flow in order to show the results element */
const showResults = () => {
    if (p()) {
        p().remove()
    }

    btn.className = "done" // Transitions the button to disappearance

    setTimeout(() => {
        btn.className = "none" // Removes the button from display
        block.className = "hidden" // Makes the div. visibly hidden; see [1]
        setTimeout(() => {
            block.className = "" // After 100ms, allows to transition up to total opacity
        }, .1 * 1000);
    }, .3 * 1000);
}

/* Inserts an element after another; taken from [2] */
const insertAfter = (newNode, referenceNode) => {
    referenceNode.parentNode.insertBefore(newNode, referenceNode.nextSibling);
}

/* Inserts an element that embeds a warning */
const createWarning = (e) => {
    const warning = document.createElement("p")
    warning.appendChild(textNode(e))
    warning.setAttribute("id", "warning"); // Adds ID to check if it has been already inserted into the page

    if (!p()) {
        insertAfter(warning, btn);
    }
}

const isURLValid = (string) => { // Robbed from [3]
	const res = string.match(/(http(s)?:\/\/.)?(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)/g);
	return (res !== null)
}

click(btn, async () => {
    if (btn.className !== "done") {
        btn.className = "clicked" // Starts the animation loop on the button

        // Checks if any text has been written inside the textarea
        const body = input.value
		if (body !== "") {
            if (p()) { // Removes the warning tag now that text has been written
                p().remove()
            }

            const splitArray = body.split("\n")
            const urls = [...new Set(splitArray.filter(n => n))] // Creates a Set, and converts it back to an array, to remove duplicates from the URLs; see [4]
                                                                 // Moreover, removes any empty elements such as empty strings; see [5]

            // Checks if any duplicates were found on the original set by comparing lengths; see [6]
            if (urls.length !== splitArray.length) {
                createWarning("Duplicate or empty text/links have been found, and removed from the query.");
            }

            const notURLError = "One of the lines doesn't contain an URL!"

            // Loops over each line, checks if it is a valid URL, and creates a link for every one
            for (const url of urls) {
				if (!isURLValid(url)) {
					showTag(notURLError); showResults(); break;
				}

				const link = document.createElement("a")
                const linkNode = textNode("Link")

                link.appendChild(linkNode);
                link.setAttribute("href", url);

                results.appendChild(link);

				await new Promise(r => setTimeout(r, 1000));
            }

            showResults(); return;
        }

        // Otherwise, it loads the animation for a while, until it loads an warning that no text has been written
        setTimeout(() => {
            btn.className = ""

            createWarning("You haven't put any URL in the text area!");
        }, 2 * 1000);
    }
});

// Reloads the page in case the button is clicked
const reload = document.getElementById("reload");

click(reload, () => {
    window.location.reload() // See [7]
})


/* Allows to auto-resize the textarea depending on the vertical length of the content; see [8] & [9] */
const autoHeight = (e) => {

    // Avoids that no unnecessary height is added to text with no more than 1 line
    const lines = e.value.split("\n").length;
    if (lines <= 1) {
        e.style.height = `auto`; return
    }

    e.style.height = `auto`; // This avoids a catastrophe (i.e. increasing the height after typing on lines next to the first one)
    e.style.height = `${e.scrollHeight}px`
}

/* Limits the amount of definite lines (spec. based on the number of line breaks) up to three; adapted from [10] & [11] */
const limitLines = (o, e) => {
    const lines = o.value.split("\n").length;
    switch(e.key) { // Uses KeyboardEvent.key instead of keyCode, or UIEvent.which as suggested by [10]; see [12] & [13]
        case "Enter":
            if (lines >= 3) {
                return false;
            }
            break;
    }
}

/* Sources:
	* [1]: https://www.impressivewebs.com/animate-display-block-none/
	* [2]: https://stackoverflow.com/a/4793630
	* [3]: https://stackoverflow.com/a/49849482
	* [4]: https://stackoverflow.com/a/9229821
	* [5]: https://stackoverflow.com/a/2843625
	* [6]: https://stackoverflow.com/a/7376645
	* [7]: https://stackoverflow.com/a/3715123
	* [8]: https://stackoverflow.com/a/60882356
	* [9]: https://youtube.com/watch?v=Yor9Y73M764
	* [10]: https://stackoverflow.com/a/557227
	* [11]: https://stackoverflow.com/a/6501310
	* [12]: https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key#examples
	* [13]: https://developer.mozilla.org/en-US/docs/Web/API/UIEvent/which
*/

function fallbackCopyTextToClipboard(text) {
    var textArea = document.createElement("textarea");
    textArea.value = text;

    // Avoid scrolling to bottom
    textArea.style.top = "0";
    textArea.style.left = "0";
    textArea.style.position = "fixed";

    document.body.appendChild(textArea);
    textArea.focus();
    textArea.select();

    try {
        var successful = document.execCommand('copy');
        document.body.removeChild(textArea);

        if (!successful) {
            throw new Error("could not copy");
        }
    } catch (err) {
        document.body.removeChild(textArea);
        throw err;
    }

}

export async function copyTextToClipboard(text) {
    if (!navigator.clipboard) {
        fallbackCopyTextToClipboard(text);
        return;
    }

    await navigator.clipboard.writeText(text);
}
const btn = document.getElementById("shortenBtn");
const input = document.getElementById("urlInput");
const resultDiv = document.getElementById("result");

btn.addEventListener("click", async () => {
  const url = input.value.trim();
  if (!url) {
    resultDiv.textContent = "Please enter a URL";
    return;
  }

  try {
    const response = await fetch("http://localhost:8080/shorten", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ url })
    });

    if (!response.ok) {
      resultDiv.textContent = "Error shortening URL";
      return;
    }

    const data = await response.json();
    resultDiv.innerHTML = `Short URL: <a href="${data.short_url}" target="_blank">${data.short_url}</a>`;
    input.value = "";
  } catch (err) {
    resultDiv.textContent = "Network error: " + err.message;
  }
});

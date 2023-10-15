const postParser = async () => {
  const markdown = document.getElementById('markdown').value;
  const lines = markdown.split("\n");
  const response = await fetch('http://localhost:8080', {
    method: 'POST',
    mode: 'no-cors',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ text: lines }),
  });

  console.log("Success:", response.body);
}
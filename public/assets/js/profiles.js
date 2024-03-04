function executeScript() {
    const selectedValue = document.getElementById('selectedValue').value;
    fetch('http://localhost:3000/executeScript', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: `choice=${selectedValue}`,
    })
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.text();
        })
        .then(() => {
            // Handle success if needed
            console.log('Script executed successfully');
        })
        .catch(error => {
            console.error('Error executing script:', error);
        });
}

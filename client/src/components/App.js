import React, { useState } from "react"
import FileUploader from './FileUploader/FileUploader'

const App = () => {
  const [name, setName] = useState("")
  const [selectedFile, setSelectedFile] = useState(null)
  
  const submitForm = () => {
    const formData = new FormData();
    formData.append("file", selectedFile);

    fetch('http://localhost:8080/file', {
      method: 'POST',
      body: formData 
    })
      .then((res) => {
        debugger
        alert("File Upload success");
      })
      .catch((err) => {
        debugger
        alert("File Upload Error")
      });
  };
  
  return (
    <div className="App">
      <form>
        <input
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />

        <FileUploader
          onFileSelect={(file) => setSelectedFile(file)}
        />

        <button onClick={submitForm}>Submit</button>
      </form>
    </div>
  )
}

export default App
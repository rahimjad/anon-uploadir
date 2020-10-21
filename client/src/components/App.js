import React, { useState } from "react"
import FileUploader from './FileUploader/FileUploader'

const App = () => {
  const [selectedFile, setSelectedFile] = useState(null)
  const [linkToFile, setLinkToFile] = useState(null)
  const [isLoading, setIsLoading] = useState(false)
  
  const handleSuccess = async res => {
    setSelectedFile(null)
    setIsLoading(false)

    const json = await res.json()

    setLinkToFile(`http://localhost:8080/file/${json.fileId}`)
  }

  const submitForm = e => {
    if (selectedFile === null) { return } 
  
    e.preventDefault()
    setIsLoading(true)

    const formData = new FormData();
    formData.append("file", selectedFile, selectedFile.name);

    fetch('http://localhost:8080/file', {
      method: 'POST',
      body: formData,
      timeout: 100000
    })
      .then((res) => {
        handleSuccess(res)
      })
      .catch((err) => {
        alert('Something is wrong')
      })
  };
  
  return (
    <div className="App">
      <form>
        <FileUploader
          onFileSelect={(file) => setSelectedFile(file)}
          disabled={isLoading || selectedFile !== null}
        />
        <button onClick={submitForm} disabled={isLoading}>Submit</button>
      </form>
      { linkToFile && 
        (<a href={linkToFile}>Click here to access your file</a>)
      }
    </div>
  )
}

export default App
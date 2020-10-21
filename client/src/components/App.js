import React, { useState } from 'react'
import styled from 'styled-components'
import FileUploader from './FileUploader/FileUploader'
import { Button, LinearProgress } from '@material-ui/core'

const StyledContentWrapper =  styled.div`
  margin: 0;
  position: absolute;
  top: 50%;
  width: 100%;
  -ms-transform: translateY(-50%);
  transform: translateY(-50%);
`
const StyledImage = styled.img`
  display: block;
  margin-left: auto;
  margin-right: auto;
  max-width: 800px;
`

const StyledForm = styled.form`
  margin-top: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
`

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
      <StyledContentWrapper>
        <StyledImage src="/title.png"/>
        <StyledForm>
          <FileUploader
            onFileSelect={(file) => setSelectedFile(file)}
            disabled={isLoading}
          />
          { selectedFile && 
            (<Button 
                size="medium" 
                onClick={submitForm} 
                disabled={isLoading} 
                color="secondary">
                  Submit To The Interwebs
              </Button>
            )
          }
        </StyledForm>
        { isLoading &&    
          (<LinearProgress color="secondary" />)
        }
        { linkToFile && 
          (<a href={linkToFile}>Click here to access your file</a>)
        }
      </StyledContentWrapper>
    </div>
  )
}

export default App
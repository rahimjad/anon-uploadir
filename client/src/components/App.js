import React, { useState } from 'react'
import styled from 'styled-components'
import FileUploader from './FileUploader/FileUploader'
import SuccessModal from './SuccessModal/SuccessModal'
import ErrorModal from './ErrorModal/ErrorModal'
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

const StyledLinearProgress = styled(LinearProgress)`
  margin-top: 50px;
`

const StyledFileName = styled.p`
  margin-top: 30px;
  color: white;
  text-align: center;
`

const App = () => {
  const [selectedFile, setSelectedFile] = useState(null)
  const [linkToFile, setLinkToFile] = useState(null)
  const [isLoading, setIsLoading] = useState(false)
  const [showSuccessModal, setShowSuccessModal] = useState(false)
  const [errorMessage, setErrorMessage] = useState(null)
  const [showErrorModal, setShowErrorModal] = useState(false)
  
  const handleResponse = async res => {
    const json = await res.json()

    if (res.status === 200) {
      handleSuccess(json.fileId)
    } else {
      handleError(json.error)
    }
  }

  const handleSuccess = fileId => {
    setIsLoading(false)
    setSelectedFile(null)

    setLinkToFile(`http://localhost:8080/file/${fileId}`)

    setShowSuccessModal(true)
  }
  
  const handleError = (errorMessage=null) => {
    setIsLoading(false)

    if (errorMessage) {
      setErrorMessage(errorMessage)
    } else {
      setErrorMessage('Whoops! Something went wrong when uploading your file.')
    }
    setShowErrorModal(true)
  }

  const submitForm = e => {
    if (selectedFile === null) { return } 

    e.preventDefault()
    setIsLoading(true)

    const formData = new FormData();
    formData.append("file", selectedFile, selectedFile.name);

    fetch('http://localhost:8080/file', {
      method: 'POST',
      body: formData
    })
      .then((res) => {
        handleResponse(res)
      })
      .catch((err) => {
        handleError()
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
                color="secondary"
              >
                  Submit To The Interwebs
              </Button>
            )
          }
        </StyledForm>
        { selectedFile &&
          (<StyledFileName>{isLoading ? 'Uploading' : 'Ready to upload'}: "{selectedFile.name}"</StyledFileName>)
        }
        { isLoading &&    
          (<StyledLinearProgress color="secondary" />)
        }
        <SuccessModal 
          open={showSuccessModal}
          handleClose={() => setShowSuccessModal(false)}
          fileLink={linkToFile}
        />  
        <ErrorModal 
          open={showErrorModal}
          handleClose={() => setShowErrorModal(false)}
          errorMessage={errorMessage}
        />  
      </StyledContentWrapper>
    </div>
  )
}

export default App
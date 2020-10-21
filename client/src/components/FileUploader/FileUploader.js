import React, { useRef } from 'react'
import { Fab } from '@material-ui/core'

const FileUploader = ({ onFileSelect, disabled }) => {
  const fileInput = useRef(null)

  const handleFileInput = (e) => {
    const file = e.target.files[0];
    onFileSelect(file);
  }

  return (
    <label htmlFor="upload-file">
      <input  
        style={{ display: 'none' }} 
        id="upload-file"  
        name="upload-file"  
        onChange={handleFileInput}
        type="file" 
      /> 

      <Fab
          color="secondary"
          size="medium"
          component="span"
          aria-label="add"
          variant="extended"
          onClick={e => fileInput.current && fileInput.current.click()} 
          disabled={disabled}
      >
          Upload File
      </Fab>
    </label>
  )
}

export default FileUploader
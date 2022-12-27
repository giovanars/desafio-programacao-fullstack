import React, { useState } from "react";

import api from './config/configApi'

function App(){

  const [file, setFile] = useState('')
  const [status, setStatus] = useState({
    type: '',
    message: ''
  });
  const [fileValid, setFileValid] = useState(true)


  const uploadFile = async e => {
    e.preventDefault();
    setStatus({
      type: '',
      message: ''
    });
    setFileValid(true);

    console.log(file)
      
    if (file.type !== "text/plain"){ 
      setFileValid(false)
      return;
    }
    
    const data = new FormData();
    data.append('file', file)
    const headers = {
      'headers': {
        'Content-Type': 'application/json'
      }
    }    

    await api.post("/upload-file", data, headers)
    .then((response) => {
      setStatus({
        type: 'success',
        message: response.data.message
      })
      alert("Upload realizado com sucesso!")
    }).catch((err) => {      
      if(err.response){
        setStatus({
          type: 'error',
          message: err.response.data.message
        })
      }else{
        setStatus({
          type: 'error',
          message: 'Ocorreu um erro, tente mais tarde'
        })
      }

    });

  }

  return(
    <div>
      <h1>Upload</h1>
      
      {status.type === 'sucess'? <p style={{color: "green"}}>{status.message}</p>: ""}
      {status.type === 'error'? <p style={{color: "red"}}>{status.message}</p>: ""}
      {fileValid ? "" : <p style={{color: "red"}}>Tipo de arquivo inválido</p>}
      
      <form onSubmit={uploadFile} >
        
        <label>Arquivo: </label>
        <input type="file" name="file"  onChange={e => setFile(e.target.files[0])}></input><br/><br/>

        
        <button type="submit">Importar</button>

        
      </form>

    </div>
  )
}

export default App;

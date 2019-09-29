import React from 'react';


const Style = {
    notFound:{
        width: '100%',
        height: '100vh',
        fontSize: '80px',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    }
};

function App() {
    return (
        <h1 style={Style.notFound}>
           login
        </h1>
    );
}

export default App;

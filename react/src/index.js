import React from 'react'
import ReactDOM from 'react-dom'
import App from './components/App'

ReactDOM.createRoot(document.getElementById('root')).render(
    <App />
);


// Router example:
{/* <BrowserRouter>
    <App>
        <Routes>
            <Route exact path='/' element={<Home/>} />
            <Route path='/about' element={<About/>} />
            <Route path='/contact' element={<Contact/>} />
            <Route path="/admin" element={<PrivateRoute><Admin /></PrivateRoute>} />
            <Route path='/login' element={<LoginPage/>} />
        </Routes>
    </App>
</BrowserRouter> */}
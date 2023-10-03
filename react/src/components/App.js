import React from 'react';
import {Routes, Route} from 'react-router-dom'

import Header from './Header.js';
import Home from '../pages/Home.js'
import Test from '../pages/Test.js'

const App = () => (
  <>
    <Header/>
    <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/test" element={<Test />} />
    </Routes>
  </>
);

export default App
// const App = ({ children }) => (
//   <>
//     <h1>ITS WORKING</h1>
//     {/* <Header /> */}
//     {children}
//     {/* <Footer /> */}
//   </>
// );

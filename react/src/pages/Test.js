import React, { useEffect, useState } from 'react';
// import axios from 'axios'
import sampleData from '../../../notes/example_movie_data.json';

// const GET_URL = ""

function Test() {
    const [data, setData] = useState([])

    useEffect(() => {
        console.log(sampleData.Results[0])
        setData(sampleData.Results[0])
        console.log(data.original_title)
    })
    // useEffect(() => {
    //     axios.get(GET_URL).then(
    //         response => {
    //             setData(response.data)
    //         }
    //     ).catch(
    //         error => console.error('Error fetching data: ', error)
    //     )
    // }, []); // The empty dependency array ensures this effect runs only once when the component is mounted)

    return (
        <div>
            <h1>Movie Data</h1>
            <strong>Title:</strong> {data["original_title"]}<br />
            <strong>Overview:</strong> {data.overview}<br />
            <strong>Rating:</strong> {data.vote_average}
        </div>
    );
}


export default Test;
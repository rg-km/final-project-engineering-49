import React from 'react';
import ReactPlayer from 'react-player';


function VideoJS() {
    return (
        <div className='videoo'>
                <ReactPlayer url='https://www.youtube.com/watch?v=kUMe1FH4CHE' controls={true}/>
        </div>
    );
}
export default VideoJS;
import React from 'react';
import ReactPlayer from 'react-player';


function VideoJS() {
    return (
        <div className='videoo'>
            
                <ReactPlayer url='https://www.youtube.com/watch?v=Ke90Tje7VS0' controls={true}/>
            
        </div>
    );
}
export default VideoJS;
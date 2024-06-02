import React from 'react';
import Login from './Login';
import parsePath from '../functions/parsePath';
import joinPath from '../functions/joinPath';
import Schema from './Schema';

export default function App() {
    // Step 1: Fetch /meta to gather type info.
    // Step 2: Fetch the data.
    // Step 3: Render the UI.
    const sessionString = localStorage.getItem('session');
    if (!sessionString) {
        return <Login />
    }
    const session = JSON.parse(sessionString);
    const path = parsePath(props.path);

    if (path.length === 0) {
        return <main><a href={"/"+session.UserID}>My Schemas</a></main>
    }
    return <Schema id={joinPath(path)} />
}

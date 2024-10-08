import React, { useState } from 'react';
import './Main.css';

const Main = () => {
  const [username, setUsername] = useState('');
  const [roomID, setRoomID] = useState('');
  const [isJoining, setIsJoining] = useState(false); // To track whether the user is joining an existing call

  const handleUsernameChange = (e) => {
    setUsername(e.target.value);
  };

  const handleRoomIDChange = (e) => {
    setRoomID(e.target.value);
  };

  const joinExistingCall = () => {
    if (username.trim() === '') {
      alert('Please enter a username before joining a call.');
    } else if (roomID.trim() === '') {
      alert('Please enter a room ID to join an existing call.');
    } else {
      alert(`Joining call as ${username} in room ${roomID}`);
      // Logic to join existing call can go here
    }
  };

  const createCall = () => {
    if (username.trim() === '') {
      alert('Please enter a username before creating a call.');
    } else {
      alert(`Creating a call as ${username}`);
      // Logic to create a new call can go here
    }
  };

  return (
    <div className="container">
      <h2>Enter your username</h2>
      <form
        onSubmit={(e) => {
          e.preventDefault();
        }}
      >
        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={handleUsernameChange}
          className="input"
        />

        {/* Conditionally render the Room ID input field if joining existing call */}
        {isJoining && (
          <input
            type="text"
            placeholder="Room ID"
            value={roomID}
            onChange={handleRoomIDChange}
            className="input"
          />
        )}

        <div className="button-container">
          <button
            type="button"
            onClick={() => setIsJoining(true)} // Show Room ID field when clicked
            className="button"
          >
            Join Existing Call
          </button>
          <button type="button" onClick={createCall} className="button">
            Create Call
          </button>
        </div>

        {/* Submit button for joining call */}
        {isJoining && (
          <div className="join-button-container">
            <button type="button" onClick={joinExistingCall} className="button">
              Join Now
            </button>
          </div>
        )}
      </form>
    </div>
  );
};

export default Main;

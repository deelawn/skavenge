import React from 'react';

/**
 * Toast notification component
 * Displays temporary notification messages to the user
 *
 * @param {object} props
 * @param {string} props.message - The message to display
 * @param {string} props.type - Type of toast: 'info', 'success', 'error'
 * @param {boolean} props.visible - Whether the toast is visible
 */
function Toast({ message, type = 'info', visible = false }) {
  return (
    <div className={`toast ${type} ${visible ? 'visible' : ''}`}>
      {message}
    </div>
  );
}

export default Toast;

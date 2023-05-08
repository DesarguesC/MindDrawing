import React from 'react';
import { useRef, useState } from 'react';
import './index.css';

function Canvas() {
  const canvasRef = useRef(null);
  const fileInputRef = useRef(null);
  const [brushSize, setBrushSize] = useState(5);
  const [canvasBackground, setCanvasBackground] = useState('');
  const [isDrawing, setIsDrawing] = useState(false);
  const [brushColor, setBrushColor] = useState('#000000');
  const [undoStack, setUndoStack] = useState([]);

  // const handleFileChange = event => {
  //   const file = event.target.files[0];
  //   const reader = new FileReader();
  //   reader.onload = event => {
  //     const imageUrl = event.target.result;
  //     setCanvasBackground(imageUrl);
  //   };
  //   reader.readAsDataURL(file);
  // };

  const handleFileChange = event => {
    const file = event.target.files[0];
    const reader = new FileReader();
    reader.onload = event => {
      const image = new Image();
      image.onload = () => {
        const canvas = canvasRef.current;
        // canvas.width = image.width;
        // canvas.height = image.height;
        const context = canvas.getContext('2d');
        context.drawImage(image, 0, 0, canvas.width, canvas.height);
      };
      image.src = event.target.result;
    };
    reader.readAsDataURL(file);
  };

  const handleMouseDown = event => {
    const canvas = canvasRef.current;
    const context = canvas.getContext('2d');
    const rect = canvas.getBoundingClientRect();
    const x = event.clientX - rect.left;
    const y = event.clientY - rect.top;
    context.beginPath();
    context.moveTo(x, y);
    setIsDrawing(true);
  };

  const handleMouseMove = event => {
    if (!isDrawing) return;
    const canvas = canvasRef.current;
    const context = canvas.getContext('2d');
    const rect = canvas.getBoundingClientRect();
    context.lineWidth = brushSize;
    const x = event.clientX - rect.left;
    const y = event.clientY - rect.top;
    context.lineTo(x, y);
    context.strokeStyle = brushColor;
    context.stroke();
  };

  const handleMouseUp = () => {
    setIsDrawing(false);
    const canvas = canvasRef.current;
    const context = canvas.getContext('2d');
    const imageData = context.getImageData(0, 0, canvas.width, canvas.height);
    setUndoStack([...undoStack, imageData]);
  };

  const handleUndo = () => {
    if (undoStack.length === 0) return;
    const canvas = canvasRef.current;
    const context = canvas.getContext('2d');
    const imageData = undoStack.pop();
    context.putImageData(imageData, 0, 0);
    setUndoStack([...undoStack]);
  };

  const handleClear = () => {
    const canvas = canvasRef.current;
    const context = canvas.getContext('2d');
    context.clearRect(0, 0, canvas.width, canvas.height);
    setUndoStack([]);
  };

  const handleSave = () => {
    const canvas = canvasRef.current;
    const link = document.createElement('a');
    link.download = 'canvas.png';
    link.href = canvas.toDataURL();
    link.click();
  };

  const handleBrushSizeChange = event => {
    setBrushSize(event.target.value);
  };

  const handleButtonClick = () => {
    fileInputRef.current.click();
  };

  return (
    <div>
      <canvas
        ref={canvasRef}
        width="800"
        height="400"
        style={{
          backgroundImage: `url(${canvasBackground})`,
          border: '1px solid black'
        }}
        onMouseDown={handleMouseDown}
        onMouseMove={handleMouseMove}
        onMouseUp={handleMouseUp}
      />
      <div>
        <button onClick={handleButtonClick}>上传图片作为背景</button>
        <input
          type="file"
          onChange={handleFileChange}
          ref={fileInputRef}
          style={{ display: 'none' }}
        />
        <div>
          <button onClick={handleUndo}>撤销</button>
          <button onClick={handleClear}>清空</button>
          <button onClick={handleSave}>保存</button>
        </div>
        <div>
          <label>画笔颜色：</label>
          <input
            type="color"
            value={brushColor}
            onChange={event => setBrushColor(event.target.value)}
          />
          <br />
          <label>画笔粗细：</label>
          <input
            type="range"
            min={1}
            max={10}
            value={brushSize}
            onChange={handleBrushSizeChange}
          />
        </div>
      </div>
    </div>
  );
}

export default Canvas;

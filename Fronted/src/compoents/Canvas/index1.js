import React, { useState, useRef, useEffect } from 'react';
import { Button, Upload } from 'antd';
import './index.css';

function Canvas() {
  const [type, setType] = useState('huabi');
  const [color, setColor] = useState('black');
  const clearRef = useRef();
  const saveRef = useRef();
  const canvasRef = useRef(null);
  const [imageUrl, setImageUrl] = useState('');
  const [canvasBackground, setCanvasBackground] = useState('');

  const handleMouseEnter = () => {
    var canvas = document.getElementById('canvas');
    const ctx = canvas.getContext('2d');

    canvas.addEventListener('mouseenter', () => {
      canvas.addEventListener('mousedown', e => {
        ctx.beginPath();
        ctx.moveTo(e.offsetX, e.offsetY);
        ctx.strokeStyle = color;
        canvas.addEventListener('mousemove', draw);
      });
      canvas.addEventListener('mouseup', () => {
        canvas.removeEventListener('mousemove', draw);
      });
    });

    function draw(e) {
      ctx.lineTo(e.offsetX, e.offsetY);
      ctx.stroke();
    }
  };

  const handleClear = () => {
    var canvas = document.getElementById('canvas');
    const ctx = canvas.getContext('2d');
    ctx.clearRect(0, 0, canvas.width, canvas.height);
  };

  const handleSave = () => {
    var canvas = document.getElementById('canvas');
    const ctx = canvas.getContext('2d');
    const compositeOperation = ctx.globalCompositeOperation;
    ctx.globalCompositeOperation = 'destination-over';
    ctx.fillStyle = '#fff';
    ctx.fillRect(0, 0, canvas.width, canvas.height);
    const imageData = canvas.toDataURL('image/png');
    ctx.putImageData(ctx.getImageData(0, 0, canvas.width, canvas.height), 0, 0);
    ctx.globalCompositeOperation = compositeOperation;
    const a = document.createElement('a');
    document.body.appendChild(a);
    a.href = imageData;
    a.download = 'myPaint';
    a.target = '_blank';
    a.click();
    document.body.removeChild(a);
  };

  const onPreview = async file => {
    console.log(file.url);
  };

  const handleUploadChange = info => {
    if (info.file.status === 'done') {
      setImageUrl(info.file.response.url);
    }
  };

  useEffect(() => {
    console.log(imageUrl);
  }, [imageUrl]);

  const handleFileChange = event => {
    const file = event.target.files[0];
    const reader = new FileReader();
    reader.onload = event => {
      const imageUrl = event.target.result;
      setCanvasBackground(imageUrl);
      const canvas = canvasRef.current;
      canvas.style.backgroundImage = `url(${imageUrl})`;
    };
    reader.readAsDataURL(file);
  };

  return (
    <>
      <div>
        <div className="toolsbar">
          <Button onClick={() => setType('huabi')}>画笔</Button>
          <Button onClick={() => setType('rect')}>正方形</Button>
          <Button onClick={() => setType('arc')}>圆形</Button>
          <span>颜色：</span>
          <input
            type="color"
            value={color}
            onChange={e => {
              setColor(e.target.value);
            }}></input>
          <Button ref={clearRef} onClick={handleClear}>
            清空
          </Button>
          <Button ref={saveRef} onClick={handleSave}>
            保存
          </Button>

          <input type="file" onChange={handleFileChange} />
        </div>
        <canvas
          ref={canvasRef}
          id="canvas"
          width="800"
          height="400"
          style={{
            border: '1px solid black',
            marginTop: '10px',
            backgroundImage: canvasBackground
          }}
          onMouseEnter={handleMouseEnter}></canvas>
      </div>
    </>
  );
}

export default Canvas;

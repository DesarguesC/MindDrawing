import React, { useEffect, useRef, useState } from 'react';
import { Checkbox, Input, Button, Col, Row, Modal } from 'antd';
import './index.less';
import {
  ArrowDown,
  Upload,
  Save,
  Back,
  Delete,
  Optimize,
  Label,
  Notes,
  Material
} from '@icon-park/react';
import '@icon-park/react/styles/index.css';
import CreatedImg from '../../assets/images/ironman2.png';
import spring from '../../assets/images/spring.png';
import summer from '../../assets/images/summer.jpg';
import autumn from '../../assets/images/autumn.png';
import winter from '../../assets/images/winter.jpg';
import { fabric } from 'fabric';

const { TextArea } = Input;

function Content() {
  const canvasRef = useRef(null);
  const fileInputRef = useRef(null);
  const fileInputRef2 = useRef(null);
  const [brushSize, setBrushSize] = useState(5);
  const [canvasBackground, setCanvasBackground] = useState('');
  const [isDrawing, setIsDrawing] = useState(false);
  const [brushColor, setBrushColor] = useState('#000000');
  const [undoStack, setUndoStack] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [isTextModalOpen, setIsTextModalOpen] = useState(false);
  const [imgUrl, setImgUrl] = useState('');
  const [text, setText] = useState('');
  const [largeImgUrl, setLargeImgUrl] = useState('');
  const [largeImgText, setLargeImgText] = useState('');

  useEffect(() => {
    if (canvasRef.current) {
      const canvas = canvasRef.current;
      canvas.width = '960';
      canvas.height = '550';
    }
  }, []);

  const handleFileChange = event => {
    const file = event.target.files[0];
    const reader = new FileReader();
    reader.onload = event => {
      const image = new Image();
      image.onload = () => {
        const canvas = canvasRef.current;
        const context = canvas.getContext('2d');
        context.drawImage(image, 0, 0, canvas.width, canvas.height);
      };
      image.src = event.target.result;
    };
    reader.readAsDataURL(file);
  };

  const handleSourceAdd = event => {
    const canvas = new fabric.Canvas('c');
    const file = event.target.files[0];
    const reader = new FileReader();
    reader.onload = event => {
      const data = event.target.result;
      fabric.Image.fromURL(data, function (img) {
        const oImg = img.set({ left: 0, top: 0, scaleX: 0.5, scaleY: 0.5 });
        canvas.add(oImg).setActiveObject(oImg);
      });
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
    const imgUrl = canvas.toDataURL();
    console.log(imgUrl);
    setImgUrl(imgUrl);
    link.href = canvas.toDataURL();
    link.click();
  };

  const handleBrushSizeChange = event => {
    setBrushSize(event.target.value);
  };

  const handleButtonClick = () => {
    fileInputRef.current.click();
  };

  const handleSourceChoose = () => {
    fileInputRef2.current.click();
  };

  const onTagChange = checkedValues => {
    console.log('checked = ', checkedValues);
  };

  const showModal = () => {
    setIsModalOpen(true);
  };

  const showTextModal = () => {
    setIsTextModalOpen(true);
    setLargeImgUrl(imageData[0].url);
    setLargeImgText(imageData[0].context);
  };

  const handleOk = () => {
    setIsModalOpen(false);
  };

  const handleTextOk = () => {
    setIsTextModalOpen(false);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
    setIsTextModalOpen(false);
  };

  const onTextChange = e => {
    setText(e.target.value);
    console.log(text);
  };

  const options = [
    { label: '卡通', value: '1' },
    { label: '山', value: '2' },
    { label: '花', value: '3' },
    { label: '天空', value: '4' },
    { label: '科技', value: '5' },
    { label: '风景', value: '6' },
    { label: '城市', value: '7' },
    { label: '建筑', value: '8' },
    { label: '森林', value: '9' },
    { label: '星空', value: '10' },
    { label: '春天', value: '11' }
  ];

  const image_description =
    '处处春光明媚，处处春意盎然。小溪的冰融化了，溪水清澈见底，小鱼在溪水里欢快的玩耍。小草偷偷的从土钻出来，伸展嫩绿的身躯。各色的花争奇斗艳，散在草丛里，像无数的星星眨呀眨。';

  const imageData = [
    {
      id: 0,
      url: spring,
      context:
        '处处春光明媚，处处春意盎然。小溪的冰融化了，溪水清澈见底，小鱼在溪水里欢快的玩耍。小草偷偷的从土钻出来，伸展嫩绿的身躯。各色的花争奇斗艳，散在草丛里，像无数的星星眨呀眨'
    },
    {
      id: 1,
      url: summer,
      context:
        '夏天是属于花的季节。各种各样的花儿这儿一朵，那儿一簇，大片大片的开出无限的美丽。红的、黄的、蓝的、粉的、白的..把夏天装扮的五彩缤纷'
    },
    {
      id: 2,
      url: autumn,
      context:
        '秋天撒下丰收的果实，天更蓝了，更纯洁，更明净了。太阳是那样明亮，亮得更加柔和。云是那样无暇，给人一种玲珑剔透的感觉，，心情也变得舒畅起来'
    },
    {
      id: 3,
      url: winter,
      context:
        '下雪的时候，一片片雪花从天上落下来，一会儿，山头白了，房子白了，窗外的一切都白了雪花落在我们的手心上化成了一滴水，洁白无瑕，晶莹剔透'
    }
  ];

  return (
    <div className="content-container">
      <div className="top">
        <i className="icon icon-xiajiantou"></i>
        <span id="huabu">
          <ArrowDown theme="outline" size="30" fill="#6e41ff" />
          画布
        </span>
        <label type="primary">画笔颜色：</label>
        <input
          style={{ marginRight: '20px' }}
          type="color"
          value={brushColor}
          onChange={event => setBrushColor(event.target.value)}
        />
        <label>画笔粗细：</label>
        <input
          style={{ marginRight: '15px' }}
          type="range"
          min={1}
          max={10}
          value={brushSize}
          onChange={handleBrushSizeChange}
        />
        <div className="buttons">
          <Button type="primary" onClick={handleButtonClick}>
            <Upload
              theme="outline"
              size="15"
              fill="white"
              style={{ marginRight: '5px' }}
            />
            上传背景图片
          </Button>
          <input
            type="file"
            onChange={handleFileChange}
            ref={fileInputRef}
            style={{ display: 'none' }}
          />
          <Button type="primary" onClick={handleSourceChoose}>
            <Material
              theme="outline"
              size="15"
              fill="white"
              style={{ marginRight: '5px' }}
            />
            选择图片素材
          </Button>
          <input
            type="file"
            // onChange={handleSourceAdd}
            onChange={() => {}}
            ref={fileInputRef2}
            style={{ display: 'none' }}
          />
          <Button type="primary" onClick={handleUndo}>
            <Back
              theme="outline"
              size="15"
              fill="white"
              style={{ marginRight: '5px' }}
            />
            撤销
          </Button>
          <Button type="primary" onClick={handleClear}>
            <Delete
              theme="outline"
              size="15"
              fill="white"
              style={{ marginRight: '5px' }}
            />
            清空
          </Button>
          <Button type="primary" onClick={handleSave}>
            <Save
              theme="outline"
              size="15"
              fill="white"
              style={{ marginRight: '5px' }}
            />
            保存
          </Button>
          <Button type="primary" onClick={showModal}>
            <Optimize
              theme="outline"
              size="15"
              fill="white"
              style={{ marginRight: '5px' }}
            />
            智慧生成绘图
          </Button>
        </div>
      </div>

      <div className="bottom">
        <div className="left">
          <canvas
            id="c"
            ref={canvasRef}
            onMouseDown={handleMouseDown}
            onMouseMove={handleMouseMove}
            onMouseUp={handleMouseUp}
          />
        </div>
        <div className="right">
          <div className="label" style={{ margin: '10px' }}>
            <div>
              <Label theme="outline" size="24" fill="#6e41ff" />
              <span>标签：</span>
            </div>
            <Checkbox.Group
              // options={options}
              defaultValue={[]}
              onChange={onTagChange}>
              <Row style={{ marginTop: '10px' }}>
                {options.map(option => {
                  return (
                    <Col span={8} key={option.label}>
                      <Checkbox value={option.value}>{option.label}</Checkbox>
                    </Col>
                  );
                })}
              </Row>
            </Checkbox.Group>
          </div>

          <div className="textarea">
            <Notes theme="outline" size="24" fill="#6e41ff" />
            <span>文本:</span>
            <TextArea
              className="text"
              onChange={e => onTextChange(e)}
              allowClear
            />
          </div>
          <div className="bb">
            <Button type="primary" onClick={showTextModal}>
              <Optimize
                theme="outline"
                size="15"
                fill="white"
                style={{ marginRight: '5px' }}
              />
              <span>智慧生成绘本</span>
            </Button>
          </div>
        </div>
      </div>
      <Modal
        title="生成图片预览"
        visible={isModalOpen}
        cancelText="放弃"
        okText="保存"
        onOk={handleOk}
        onCancel={handleCancel}>
        <img src={CreatedImg} alt="tp" width="100%" height="100%" />
      </Modal>
      <Modal
        title="生成绘本预览"
        visible={isTextModalOpen}
        cancelText="放弃"
        okText="保存"
        onOk={handleTextOk}
        onCancel={handleCancel}>
        <div id="image-container">
          <div id="large-image">
            <img src={largeImgUrl}></img>
            <div className="img_text">{largeImgText}</div>
          </div>
          <div className="small-images">
            <ul>
              <li>
                {imageData.map(item => {
                  return (
                    <div
                      key={item.id}
                      onClick={() => {
                        setLargeImgText(imageData[item.id].context);
                        setLargeImgUrl(imageData[item.id].url);
                      }}>
                      <img src={item.url} />
                    </div>
                  );
                })}
              </li>
            </ul>
          </div>
        </div>
      </Modal>
    </div>
  );
}

export default Content;

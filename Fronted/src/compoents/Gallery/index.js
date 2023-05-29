import { Card } from 'antd';
import React from 'react';
import './index.less';

const data = [
  {
    title: 'example',
    url: 'https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png',
    description: 'www.instagram.com'
  },
  {
    title: '2',
    url: 'https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png',
    description: 'hh'
  },
  {
    title: '3',
    url: 'https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png',
    description: 'hh'
  },
  {
    title: '4',
    url: 'https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png',
    description: 'hh'
  },
  {
    title: '5',
    url: 'https://os.alipayobjects.com/rmsportal/QBnOOoLaAfKPirc.png',
    description: 'hh'
  }
];

const { Meta } = Card;
const App = () => (
  <div className="container">
    <div className="grid">
      {data.map(data => {
        return (
          <div
            className="card"
            key={data.title}
            onClick={() => console.log(data.title)}>
            <img src={data.url} alt={data.title} />
            <div>{data.title}</div>
            <div>{data.description}</div>
          </div>
        );
      })}
    </div>
  </div>
);

export default App;

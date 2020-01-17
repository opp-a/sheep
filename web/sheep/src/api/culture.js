// import axios from 'axios'

export const getCulturePictures = data => {
  // const pageNumber = data.pageNumber
  const pictures = []
  if (data.page > 3) {
    return pictures
  }
  return [
    {
      time: '2019-12-19',
      pictures: [
        {
          name: 'picture-1',
          src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
          time: '2019-12-19'
        },
        {
          name: 'picture-2',
          src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
          price: '2019-12-19'
        },
        {
          name: 'picture-3',
          src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
          price: '2019-12-19'
        }
      ]
    },
    {
      time: '2019-12-17',
      pictures: [
        {
          name: 'picture-4',
          src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
          price: '2019-12-17'
        },
        {
          name: 'picture-5',
          src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
          price: '2019-12-17'
        }
      ]
    },
    {
      time: '2019-12-16',
      pictures: [
        {
          name: 'picture-6',
          src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
          price: '2019-12-16'
        },
        {
          name: 'picture-7',
          src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
          price: '2019-12-16'
        },
        {
          name: 'picture-8',
          src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
          price: '2019-12-16'
        },
        {
          name: 'picture-9',
          src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
          price: '2019-12-16'
        },
        {
          name: 'picture-10',
          src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
          price: '2019-12-16'
        },
        {
          name: 'picture-11',
          src: 'https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png',
          price: '2019-12-16'
        }
      ]
    }
  ]
}

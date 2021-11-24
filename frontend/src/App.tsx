import React, {useState} from 'react';
import useAxios from 'axios-hooks';

const App = () => {
  const [name, setName] = useState('');

  const [{data: getData, error: getError}] = useAxios(
    `${process.env.REACT_APP_API_ENDPOINT}/your-name`
  );

  const [{data: postData, error: postError}, sendName] = useAxios(
    {
      url: `${process.env.REACT_APP_API_ENDPOINT}/your-name`,
      method: 'post'
    },
    {
      manual: true,
    });

  const onSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const data = new FormData();
    data.set('name', name);
    sendName({data});
    setName('');
  }

  return (
    <>
      {getData?.message || postData?.message || getError?.response?.data?.message || postError?.response?.data.message}
      <form onSubmit={onSubmit}>
        <input type="text" value={name} onChange={e => setName(e.target.value)}/>
        <input type="submit"/>
      </form>
    </>
  )
}

export default App

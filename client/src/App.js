import React from "react";
import API from "./utils/API";
import Animal from "./Animal";

class App extends React.Component{
  constructor(props){
    super(props);

    this.state = {
      isLoading: true,
      id: null,
      avatar: null,
      email: null
    };
  }
  render(){

    const { id, name, age, isLoading } = this.state;

    return (<Animal isLoading={isLoading} id={id} name={name} age={age} />);
  }

  async componentDidMount(){
    // Load async data.
    let animalData = await API.get('/', {
      params:{
        results:1,
        inc: 'id,name,age'
      }
    });
    animalData = animalData.data.results[0];
    const id = animalData.id;
    const name = animalData.name;
    const age = animalData.age;

    this.setState({
      ...this.state, ...{
        isLoading: false,
        id,
        name,
        age
      }
    });
    // Update state with new data.
    // Re-render our component.
  }

}

export default App;
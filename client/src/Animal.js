import React from "react";
import PropTypes from "prop-types";
import "bootstrap/dist/css/bootstrap.min.css";
import "shards-ui/dist/css/shards.min.css";
import { Card } from "shards-react";
import { Button } from "shards-react";
class User extends React.Component {
  render() {
    const { id, name, age, isLoading } = this.props;
const animalDetails = (
      <div>
        <h3>{id}</h3>
        <h4 className="mb-0">{name}</h4>
        <span className="text-muted">{age}</span>
      </div>
    );
const loadingMessage = <span className="d-flex m-auto">Loading...</span>;
return (
      <Card
        className="mx-auto mt-4 text-center p-5"
        style={{ maxWidth: "300px", minHeight: "250px", backgroundColor: "rgba(255, 255, 255, 0.94)"}}
      >
        {isLoading ? loadingMessage : animalDetails}
      </Card>
    );
  }
}
User.propTypes = {
  id: PropTypes.string,
  name: PropTypes.string,
  age: PropTypes.string,
  isLoading: PropTypes.bool
};
export default User;
import React, {Component} from 'react';

class CreateAccount extends Component {
    constructor(props) {
        super(props);
        this.state = {
            email: "",
            name: ""
        }

    }

    changeEmail = (event) => {
        this.setState({
           email: event.target.value
        })
    };

    changePin = (event) => {
        this.setState( {
            pin: parseInt(event.target.value)
        })
    }

    changeName = (event) => {
        this.setState( {
            name: event.target.value
        })
    }

    render() {
        return (
            <div>
                <p>Email:</p>
                <textarea
                    rows={1}
                    type="string"
                    value={this.state.email}
                    onChange={this.changeEmail}
                />
                <br></br>
                <p>Pin (4 digit number):</p>
                <textarea
                    rows={1}
                    type="number"
                    min={1}
                    max={9999}
                    value={this.state.pin}
                    onChange={this.changePin}
                />
                <br></br>
                <p>Name:</p>
                <textarea
                    rows={1}
                    type="string"
                    value={this.state.name}
                    onChange={this.changeName}
                />
            </div>
        )
    }
}
export default CreateAccount;

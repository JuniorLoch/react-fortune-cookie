import React, { Component } from "react";
import "./Botaobiscoito.css";

class Botaobiscoito extends Component {
    constructor(props) {
        super(props);
        this.state = {};
    }

    render() {
        return (
            <div id="biscoito">
                <button onClick={this.props.clicarBotao}>
                    <img alt="" src={require("../../Assets/biscoito.png")} />
                </button>
            </div>
        );
    }
}

export default Botaobiscoito;

import React, { Component } from 'react';
import Botaobiscoito from './Components/Botaobiscoito/Botaobiscoito';
import Camposorte from './Components/Camposorte/Camposorte';
import api from './services/api';
import translate from 'translate';
import './App.css';

class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            msg: '',
            anim: false,
        };

        // this.frases = [
        //     "Siga os bons e aprenda com eles.",
        //     "O bom-senso vale mais do que muito conhecimento.",
        //     "O riso é a menor distância entre duas pessoas.",
        //     "Deixe de lado as preocupações e seja feliz.",
        //     "Realize o óbvio, pense no improvável e conquiste o impossível.",
        //     "Acredite em milagres, mas não dependa deles.",
        //     "A maior barreira para o sucesso é o medo do fracasso.",
        // ];

        this.clicarBotao = this.clicarBotao.bind(this);
        this.msgAleatoria = this.msgAleatoria.bind(this);
        //this.controleAnim = this.controleAnim.bind(this);
    }

    async clicarBotao() {
        // this.controleAnim(true);
        this.setState({
            msg: await translate(await this.msgAleatoria(), 'Portuguese'),
        });
    }

    async msgAleatoria() {
        // var rand = Math.random() * this.frases.length;
        // rand = Math.floor(rand);
        // var randmsg = this.frases[rand];
        // // eslint-disable-next-line
        // while (this.state.msg == randmsg) {
        //     rand = Math.random() * this.frases.length;
        //     rand = Math.floor(rand);
        //     randmsg = this.frases[rand];
        // }
        var randmsg = await api.get('/fortunes');

        randmsg = randmsg.data;

        console.log(randmsg);
        return randmsg ? randmsg : 'nao encontrado';
    }

    // controleAnim(result) {
    //     this.setState({ anim: result });
    // }

    render() {
        return (
            <div id="conteiner">
                <p id="texto">Sorte do dia</p>
                <Botaobiscoito clicarBotao={this.clicarBotao} />
                <Camposorte msg={this.state.msg} />
            </div>
        );
    }
}

export default App;

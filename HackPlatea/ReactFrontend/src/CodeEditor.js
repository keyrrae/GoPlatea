/**
 * Created by xuanwang on 5/9/17.
 */
import React from 'react';
import { Editor, EditorState} from 'draft-js';
import axios from 'axios'

class CodeEditor extends React.Component {
    constructor(props) {
        super(props);
        this.state = {editorState: EditorState.createEmpty(), text: ""};
        this.focus = () => this.refs.editor.focus();
        this.onChange = (editorState) => this.setState({editorState});
        this.logState = () => {console.log((this.state.editorState.getCurrentContent().getPlainText())); this.setState({text:"log"});};
        this.runCode = () =>{
            axios.post('http://localhost:8000/', this.state.editorState.getCurrentContent().getPlainText())
                .then((response) => this.setState({ text: response.data }));
        }
    }
    render() {
        return (
            <div style={styles.root}>
                <div style={styles.editor} onClick={this.focus}>
                    <Editor
                        editorState={this.state.editorState}
                        onChange={this.onChange}
                        //placeholder="Enter some text..."
                        ref="editor"
                    />
                </div>
                <input
                    onClick={this.logState}
                    style={styles.button}
                    type="button"
                    value="Log State"
                />
                <input
                    onClick={this.runCode}
                    style={styles.button}
                    type="button"
                    value="Run Code"
                />
                <input
                    onClick={this.logState}
                    style={styles.button}
                    type="button"
                    value="About"
                />
                <ResultText text={this.state.text}/>
            </div>
        );
    }
}

class ResultText extends React.Component {
    render(){
        return (
            <div>
                <p>
                    {this.props.text}
                </p>
            </div>
        );
    }

}

const styles = {
    root: {
        fontFamily: '\'Helvetica\', sans-serif',
        padding: 20,
        width: 600,
    },
    editor: {
        border: '1px solid #ccc',
        cursor: 'text',
        minHeight: 80,
        padding: 10,
    },
    button: {
        marginTop: 10,
        textAlign: 'center',
    },
};

export default CodeEditor;
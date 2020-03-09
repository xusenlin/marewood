import React from 'react';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import DialogTitle from '@material-ui/core/DialogTitle';
import Dialog from '@material-ui/core/Dialog';


export default function SwitchBranchDialog(props) {
    const [branchList, setBranchList] = React.useState(["master"]);
    const { onClose, selectedValue,jobId,...other } = props;

    function handleClose() {
        onClose(selectedValue);
    }

    function handleListItemClick(value) {
        onClose(value);
    }

    return (
        <Dialog onClose={handleClose} aria-labelledby="dialog-title" {...other}>
            <DialogTitle id="dialog-title">切换当前任务的分支</DialogTitle>
            <List>
                {branchList.map(branch => (
                    <ListItem button onClick={() => handleListItemClick(branch)} key={branch}>
                        <ListItemText primary={branch} />
                    </ListItem>
                ))}
            </List>
        </Dialog>
    );
}

SwitchBranchDialog.propTypes = {
    onClose: PropTypes.func,
    open: PropTypes.bool,
    jobId: PropTypes.number,
    selectedValue: PropTypes.string,
};

// export default function SimpleDialogDemo() {
//     const [open, setOpen] = React.useState(false);
//     const [selectedValue, setSelectedValue] = React.useState(emails[1]);
//
//     function handleClickOpen() {
//         setOpen(true);
//     }
//
//     const handleClose = value => {
//         setOpen(false);
//         setSelectedValue(value);
//     };
//
//     return (
//         <div>
//             <Typography variant="subtitle1">Selected: {selectedValue}</Typography>
//             <br />
//             <Button variant="outlined" color="primary" onClick={handleClickOpen}>
//                 Open simple dialog
//             </Button>
//             <SimpleDialog selectedValue={selectedValue} open={open} onClose={handleClose} />
//         </div>
//     );
// }
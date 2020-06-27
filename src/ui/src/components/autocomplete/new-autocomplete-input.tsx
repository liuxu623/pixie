import clsx from 'clsx';
import * as React from 'react';

import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';

import { Key } from './key';

const useStyles = makeStyles((theme: Theme) => createStyles({
  root: {
    ...theme.typography.h6,
    cursor: 'text',
    padding: theme.spacing(2.5),
    fontWeight: theme.typography.fontWeightLight,
    display: 'flex',
    flexDirection: 'row',
  },
  inputElem: {
    flex: 1,
    ...theme.typography.h6,
    position: 'absolute',
    opacity: 0,
    width: '100%',
    border: 0,
    fontWeight: theme.typography.fontWeightLight,
  },
  fixedInput: {
    width: '100%',
    overflow: 'hidden',
  },
  caret: {
    display: 'inline-block',
    width: 0,
    height: 'auto',
    borderRight: '1px solid white',
    visibility: 'hidden',
    '&.visible': {
      visibility: 'visible',
    },
    position: 'absolute',
  },
  inputValue: {
    flex: 1,
    whiteSpace: 'nowrap',
    display: 'inline-block',
    position: 'relative',
  },
  prefix: {
    paddingRight: theme.spacing(2),
  },
  inputKey: {
    color: `${theme.palette.success.main}`,
  },
  hint: {
    opacity: 0.2,
  },
}));

export interface AutocompleteField {
  type: 'key' | 'value'; // The type signifies what color the text should be.
  value: string;
}

interface AutocompleteInputProps {
  onChange: (val: string, cursor: number) => void;
  onKey: (key: Key) => void;
  className?: string;
  value: Array<AutocompleteField>; // The text to show in the input box, marked by type.
  cursorPos: number; // Where the cursor position should be.
  prefix?: React.ReactNode; // An image to display to the left of the input box, if any.
  // Clicking around in the input box may result in a cursor change. This allows AutocompleteInput to notify the
  // parentClass of this change.
  setCursor: (val: number) => void;
  placeholder?: string; // The text to show if the input box is empty.
}

export const AutocompleteInput: React.FC<AutocompleteInputProps> = ({
  onChange,
  onKey,
  className,
  cursorPos,
  setCursor,
  prefix = null,
  placeholder = '',
  value,
}) => {
  const classes = useStyles();
  const [focused, setFocused] = React.useState<boolean>(true);
  const inputRef = React.useRef<HTMLInputElement>(null);

  const handleChange = React.useCallback((e) => {
    onChange(e.target.value, e.target.selectionStart);
  }, [onChange]);

  const handleKey = React.useCallback((e) => {
    switch (e.key) {
      case 'Tab':
        e.preventDefault();
        onKey('TAB');
        break;
      case 'Enter':
        onKey('ENTER');
        break;
      case 'ArrowUp':
        onKey('UP');
        break;
      case 'ArrowDown':
        onKey('DOWN');
        break;
      case 'ArrowRight':
        onKey('RIGHT');
        break;
      case 'ArrowLeft':
        onKey('LEFT');
        break;
      case 'Backspace':
        onKey('BACKSPACE');
        e.preventDefault();
        break;
      default:
        e.target.setSelectionRange(cursorPos, cursorPos);
    }
  }, [onKey, cursorPos]);

  const handleFocus = React.useCallback(() => {
    setFocused(true);
  }, []);

  const handleBlur = React.useCallback(() => {
    setFocused(false);
  }, []);

  const focusInput = React.useCallback(() => {
    // The user has clicked a new position in the input box. We must
    // update the cursor position accordingly.
    setCursor(inputRef.current.selectionStart);
    inputRef.current.focus();
  }, []);

  // Focus the input element whenever the suggestion changes.
  React.useEffect(() => {
    inputRef.current.focus();
  }, [value]);

  // Each field in the input box is a different span so that we can control color.
  const fieldSpans = [];
  let displayText = '';
  value.forEach((v, index) => {
    const className = v.type === 'value' ? '' : classes.inputKey;
    if (cursorPos >= displayText.length && cursorPos <= displayText.length + v.value.length) {
      // The cursor should be displayed in this field.
      // We need to split this field into two spans so we can add the cursor in between.
      const fieldCursor = cursorPos - displayText.length;
      fieldSpans.push(<span key={`${index}-1`} className={className}>{v.value.substring(0, fieldCursor)}</span>);
      fieldSpans.push(<Caret key={`${index}-2`} active={focused} />);
      fieldSpans.push(<span key={`${index}-caret`} className={className}>{v.value.substring(fieldCursor)}</span>);
    } else {
      fieldSpans.push(<span key={index} className={className}>{v.value}</span>);
    }
    displayText += v.value;
  });

  return (
    <div className={clsx(classes.root, className)} onClick={focusInput}>
      {prefix && <div className={classes.prefix}>{prefix}</div>}
      <div className={classes.fixedInput}>
        <div className={classes.inputValue}>
          <input
            className={classes.inputElem}
            ref={inputRef}
            value={displayText}
            onChange={handleChange}
            onFocus={handleFocus}
            onBlur={handleBlur}
            onKeyDown={handleKey}
          />
          {fieldSpans}
        </div>
        <span className={classes.hint} tabIndex={-1}>{fieldSpans.length === 0 ? placeholder : ''}</span>
      </div>
    </div>
  );
};

const BLINK_INTERVAL = 500; // 500ms = .5s

const Caret: React.FC<{ active: boolean }> = ({ active }) => {
  const classes = useStyles();
  const [visible, setVisible] = React.useState(true);
  React.useEffect(() => {
    if (!active) { return; }

    const intervalSub = setInterval(() => {
      setVisible((show) => !show);
    }, BLINK_INTERVAL);

    return () => {
      clearInterval(intervalSub);
    };
  }, [active]);
  return (
    <div className={clsx(classes.caret, active && visible && 'visible')}>&nbsp;</div>);
};

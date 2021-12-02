package screen

import (
	"RenG/src/config"
	"RenG/src/core"
	"RenG/src/lang/ast"
	"RenG/src/lang/object"
	"fmt"
	"sync"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

var (
	ScreenMutex = &sync.RWMutex{}
)

func isInTexture(texture *core.SDL_Texture, x, y int) bool {
	return x >= core.ResizeInt(config.Width, config.ChangeWidth, texture.Xpos) &&
		x <= core.ResizeInt(config.Width, config.ChangeWidth, texture.Width)+core.ResizeInt(config.Width, config.ChangeWidth, texture.Xpos) &&
		y >= core.ResizeInt(config.Height, config.ChangeHeight, texture.Ypos) &&
		y <= core.ResizeInt(config.Height, config.ChangeHeight, texture.Height)+core.ResizeInt(config.Height, config.ChangeHeight, texture.Ypos)
}

func isFirstPriority(name string) bool {
	if len(config.ScreenPriority) <= 0 {
		return false
	}
	return config.ScreenPriority[len(config.ScreenPriority)-1] == name
}

func FindScreenPriority(name string) int {
	for i := 0; i < len(config.ScreenPriority); i++ {
		if config.ScreenPriority[i] == name {
			return i
		}
	}

	return -1
}

func isScreenEnd(name string) bool {
	_, ok := config.ScreenAllIndex[name]
	return !ok
}

func isKeyWant(keyName string, inputKey uint8) bool {
	switch keyName {
	case "0":
		return inputKey == uint8(core.SDLK_0)
	case "1":
		return inputKey == uint8(core.SDLK_1)
	case "2":
		return inputKey == uint8(core.SDLK_2)
	case "3":
		return inputKey == uint8(core.SDLK_3)
	case "4":
		return inputKey == uint8(core.SDLK_4)
	case "5":
		return inputKey == uint8(core.SDLK_5)
	case "6":
		return inputKey == uint8(core.SDLK_6)
	case "7":
		return inputKey == uint8(core.SDLK_7)
	case "8":
		return inputKey == uint8(core.SDLK_8)
	case "9":
		return inputKey == uint8(core.SDLK_9)
	case "A":
		return inputKey == uint8(core.SDLK_a)
	case "B":
		return inputKey == uint8(core.SDLK_b)
	case "C":
		return inputKey == uint8(core.SDLK_c)
	case "D":
		return inputKey == uint8(core.SDLK_d)
	case "E":
		return inputKey == uint8(core.SDLK_e)
	case "F":
		return inputKey == uint8(core.SDLK_f)
	case "G":
		return inputKey == uint8(core.SDLK_g)
	case "H":
		return inputKey == uint8(core.SDLK_h)
	case "I":
		return inputKey == uint8(core.SDLK_i)
	case "J":
		return inputKey == uint8(core.SDLK_j)
	case "K":
		return inputKey == uint8(core.SDLK_k)
	case "L":
		return inputKey == uint8(core.SDLK_l)
	case "M":
		return inputKey == uint8(core.SDLK_m)
	case "N":
		return inputKey == uint8(core.SDLK_n)
	case "O":
		return inputKey == uint8(core.SDLK_o)
	case "P":
		return inputKey == uint8(core.SDLK_p)
	case "Q":
		return inputKey == uint8(core.SDLK_q)
	case "R":
		return inputKey == uint8(core.SDLK_r)
	case "S":
		return inputKey == uint8(core.SDLK_s)
	case "T":
		return inputKey == uint8(core.SDLK_t)
	case "U":
		return inputKey == uint8(core.SDLK_u)
	case "V":
		return inputKey == uint8(core.SDLK_v)
	case "W":
		return inputKey == uint8(core.SDLK_w)
	case "X":
		return inputKey == uint8(core.SDLK_x)
	case "Y":
		return inputKey == uint8(core.SDLK_y)
	case "Z":
		return inputKey == uint8(core.SDLK_z)
	case "a":
		return inputKey == uint8(core.SDLK_a)
	case "b":
		return inputKey == uint8(core.SDLK_b)
	case "c":
		return inputKey == uint8(core.SDLK_c)
	case "d":
		return inputKey == uint8(core.SDLK_d)
	case "e":
		return inputKey == uint8(core.SDLK_e)
	case "f":
		return inputKey == uint8(core.SDLK_f)
	case "g":
		return inputKey == uint8(core.SDLK_g)
	case "h":
		return inputKey == uint8(core.SDLK_h)
	case "i":
		return inputKey == uint8(core.SDLK_i)
	case "j":
		return inputKey == uint8(core.SDLK_j)
	case "k":
		return inputKey == uint8(core.SDLK_k)
	case "l":
		return inputKey == uint8(core.SDLK_l)
	case "m":
		return inputKey == uint8(core.SDLK_m)
	case "n":
		return inputKey == uint8(core.SDLK_n)
	case "o":
		return inputKey == uint8(core.SDLK_o)
	case "p":
		return inputKey == uint8(core.SDLK_p)
	case "q":
		return inputKey == uint8(core.SDLK_q)
	case "r":
		return inputKey == uint8(core.SDLK_r)
	case "s":
		return inputKey == uint8(core.SDLK_s)
	case "t":
		return inputKey == uint8(core.SDLK_t)
	case "u":
		return inputKey == uint8(core.SDLK_u)
	case "v":
		return inputKey == uint8(core.SDLK_v)
	case "w":
		return inputKey == uint8(core.SDLK_w)
	case "y":
		return inputKey == uint8(core.SDLK_y)
	case "z":
		return inputKey == uint8(core.SDLK_z)
	default:
		return false
	}
}

func applyFunction(fn object.Object, args []object.Object, name string) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := ScreenEval(fn.Body, extendedEnv, name)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return newError("not a function: %s", fn.Type())
	}
}

func extendFunctionEnv(def *object.Function, args []object.Object) *object.Environment {
	env := object.NewEncloseEnvironment(def.Env)

	for paramIdx, param := range def.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

func isCurrentExp(index int, str *ast.StringLiteral) bool {
	for i := 0; i < len(str.Exp); i++ {
		if index == str.Exp[i].Index {
			return true
		} else if index < str.Exp[i].Index {
			return false
		}
	}
	return false
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func isAssign(operator string) bool {
	switch operator {
	case "=":
		return true
	case "+=":
		return true
	case "-=":
		return true
	case "*=":
		return true
	case "/=":
		return true
	case "%=":
		return true
	default:
		return false
	}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

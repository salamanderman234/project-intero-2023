<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\User;
use Illuminate\Support\Facades\Hash;
use Illuminate\Support\Facades\Validator;

class AuthController extends Controller
{
    /**
     * Create a new AuthController instance.
     *
     * @return void
     */
    // public function __construct()
    // {
    //     $this->middleware('auth:api', ['except' => ['login','register']]);
    // }

    public function deleteUser(User $user) {
        $user->delete();
        return response()->json(["message" => "ok"],200);
    }
    public function changePassword(Request $request, User $user) {
        if($request->has("password")) {
            $password = $request->input('password');
            $hashed = Hash::make($password);

            $user->password = $hashed;
            $user->save();
        }
        return response()->json(["message" => "ok"], 200);
    }
    public function getUser(User $user) {
        return response()->json([
            "message" => "ok",
            "data" => $user,
        ],200);
    }

    public function register()
    {
        $validator = Validator::make(request()->all(),[
            'name' =>'required',
            'email' =>'required|unique:users,email',
            'password' => 'required',
            'role' =>'required|in:admin,teacher,student',
        ]);

        if ($validator->fails()) {
            return response()->json(['error' => $validator->errors()], 422);
        }

        $user = User::create([
            'name' => request('name'),
            'email' => request('email'),
            'password' => Hash::make(request('password')),
            'role' => request('role')
        ]);

        if ($user){
            return response()->json(['message' => 'Successfully Registered', "data" => $user]);
        }else{
            return response()->json(['message' => 'Something went wrong']);
        }
    }
    /**
     * Get a JWT via given credentials.
     *
     * @return \Illuminate\Http\JsonResponse
     */
    public function login()
    {
        $credentials = request(['email', 'password']);

        if (! $token = auth()->attempt($credentials)) {
            return response()->json(['error' => 'Unauthorized'], 401);
        }

        return $this->respondWithToken($token);
    }

    /**
     * Get the authenticated User.
     *
     * @return \Illuminate\Http\JsonResponse
     */
    public function me()
    {
        $user = auth()->user();
        if (empty($user)) {
            return response()->json(['error' => 'Unauthorized'], 401);
        }
        return response()->json($user);
    }

    /**
     * Log the user out (Invalidate the token).
     *
     * @return \Illuminate\Http\JsonResponse
     */
    public function logout()
    {
        auth()->logout();

        return response()->json(['message' => 'Successfully Logged Out']);
    }

    /**
     * Refresh a token.
     *
     * @return \Illuminate\Http\JsonResponse
     */
    public function refresh()
    {
        return $this->respondWithToken(auth()->refresh());
    }

    /**
     * Get the token array structure.
     *
     * @param  string $token
     *
     * @return \Illuminate\Http\JsonResponse
     */
    protected function respondWithToken($token)
    {
        return response()->json([
            'access_token' => $token,
            'token_type' => 'bearer',
            'expires_in' => auth()->factory()->getTTL() * 60
        ]);
    }
}

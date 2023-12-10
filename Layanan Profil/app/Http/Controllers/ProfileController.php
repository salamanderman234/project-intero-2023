<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Validator; 
use Illuminate\Support\Str;
use App\Models\Student;
use App\Models\Teacher;
use App\Models\Subject;

class ProfileController extends Controller
{
    public function getSiswaInfo(Request $request)
    {
        // Get the current login siswa
        // $user = Auth::user();

        // For testing
        $user = \App\Models\User::find(2);

        if ($user) {
            // Get the associated student record
            $student = $user->student;

            // Check if the associated student record exists
            if ($student) {
                // Get the data that can be exposed
                $userData = [
                    'name' => $student->name,
                    'student_number' => $student->student_number,
                    'place_of_birth' => $student->place_of_birth,
                    'date_of_birth' => $student->date_of_birth,
                    'address' => $student->address,
                    'no_handphone' => $student->no_handphone,
                    'profile_pic' => $student->profile_pic,
                ];

                return response()->json([
                    'message' => 'Data retrieved successfully',
                    'data' => $userData,
                ], 200);
            } else {
                return response()->json([
                    'message' => 'Student record not found',
                ], 404);
            }
        } else {
            return response()->json([
                'message' => 'User not found',
            ], 404);
        }
    }

    public function getTeacherInfo(Request $request)
    {
        // get the current login guru
        // $user = Auth::user();

        // for testing
        $user = \App\Models\User::find(1);

        if ($user) {
            // Get the associated teacher record
            $teacher = $user->teacher;

            // Check if the associated student record exists
            if($teacher) {
                // get the data that can be exposed
                $userData = [
                    'name' => $teacher->name,
                    'subjects' => $teacher->subject->name,
                    'date_of_birth' => $teacher->date_of_birth,
                    'address' => $teacher->address,
                    'no_handphone' => $teacher->no_handphone,
                    'email' => $teacher->email,
                    'employee_number' => $teacher->employee_number,
                ];

                return response()->json([
                    'message' => 'Data retrieved successfully',
                    'data' => $userData,
                ], 200);
            } else {
                return response()->json([
                    'message' => 'Teacher record not found'
                ], 404);
            }
        } else {
            return response()->json([
                'message' => 'User not found',
            ], 404);
        }
    }

    public function updateSiswa(Request $request)
    {
        // get the login user
        // $user = Auth::user();

        // for testing
        $user = \App\Models\User::find(2);

        if ($user) {
            // get the associated student record
            $student = $user->student;

            //check if the associated student record exist
            if($student) {
                // validasi input
                $validator = Validator::make($request->all(), [
                    'name' => 'required',
                    //'student_number' => 'required', (student number cant be edited)
                    'place_of_birth' => 'required',
                    'date_of_birth' => 'required',
                    'address' => 'required',
                    'no_handphone' => 'required',
                    'profile_pic' => 'required',
                ], [
                    'name.required' => 'Your name is required',
                    'place_of_birth.required' => 'Your place of birth is required',
                    'date_of_birth.required' => 'Your date of birth is required',
                    'address.required' => 'Your address is required',
                    'no_handphone.required' => 'Your phone number is required',
                    'profile_pic.required' => 'Your profile picture is required',
                ]);

                $errorResponse = [
                    "error" => $validator->errors(),
                ];

                $successResponse = [
                    "message" => "Your profile was updated"
                ];

                if ($validator->fails()) {
                    return response()->json([$errorResponse], 400);
                } else {
                    $student->name = $request->input('name');
                    $student->place_of_birth = $request->input('place_of_birth');
                    $student->date_of_birth = $request->input('date_of_birth');
                    $student->address = $request->input('address');
                    $student->no_handphone = $request->input('no_handphone');
                    $student->profile_pic = $request->input('profile_pic');
                    $student->save();

                    return response()->json($successResponse, 200);
                }
            } else {
                return response()->json([
                    'message' => 'Student record not found'
                ], 404);
            }
        } else {
            return response()->json([
                'message' => 'User not found',
            ], 404);
        }
    }

    public function updateTeacher(Request $request)
    {
        // get the login user
        // $user = Auth::user();

        // for testing only
        $user = \App\Models\User::find(1);

        if($user) {
            // get the associated teacher record
            $teacher = $user->teacher;

            // check if the associated teacher record exist
            if($teacher) {
                // validasi
                $validator = Validator::make($request->all(), [
                    'name' => 'required',
                    'subjects'=> 'required',
                    'date_of_birth' => 'required',
                    'address' => 'required',
                    'no_handphone' => 'required',
                    'email' => 'required',
                    'employee_number' => 'required',
                ], [
                    'name.required' => 'Your name is required',
                    'subjects.required' => 'Subjects required',
                    'date_of_birth.required' => 'Your date of birth is required',
                    'address.required' => 'Your address is required',
                    'no_handphone.required' => 'Your phone number is required',
                    'email.required' => 'Your email is required',
                ]);
                
                $errorResponse = [
                    "error" => $validator->errors(),
                ];

                $successResponse = [
                    "message" => "Your profile was updated"
                ];
                
                if ($validator->fails()) {
                    return response()->json([$errorResponse], 400);
                } else {
                    $teacher->name = $request->input('name');
                    $teacher->date_of_birth = $request->input('date_of_birth');
                    $teacher->address = $request->input('address');
                    $teacher->no_handphone = $request->input('no_handphone');
                    $teacher->email = $request->input('email');
                    $teacher->updated_at = date("H:i:s", time());
                    $teacher->save();

                    return response()->json($successResponse, 200);
                }
            } else {
                return response()->json([
                    'message' => "Teacher record not found"
                ], 404);
            }
        } else {
            return response()->json([
                'message' => 'User not found',
            ], 404);
        }
    }

    public function createSiswa(Request $request)
    {
        // Get the current login siswa
        // $user = Auth::User();

        // for testing only
        $user = \App\Models\User::find(2);

        if($user) {
            $validator = Validator::make($request->all(), [
                'name' => 'required',
                'place_of_birth' => 'required',
                'date_of_birth' => 'required',
                'address' => 'required',
                'no_handphone' => 'required',
                'profile_pic' => 'required',
            ], [
                'name.required' => 'Your name is required',
                'place_of_birth.required' => 'Your place of birth is required',
                'date_of_birth.required' => 'Your date of birth is required',
                'address.required' => 'Your address is required',
                'no_handphone.required' => 'Your phone number is required',
                'profile_pic.required' => 'Your profile picture is required',
            ]);

            $errorResponse = [
                "error" => $validator->errors(),
            ];

            $successResponse = [
                "message" => "Your profile was created"
            ];

            if ($validator->fails()) {
                return response()->json([$errorResponse], 400);
            } else {

                // if user already has a profile
                if (Student::where('user_id', $user->id)->exists()) {
                    return response()->json(['error' => 'User already has a student profile'], 400);
                }
                $student = new Student();

                // generating random 5 number for student number and making sure its unique
                do {
                    $studentNumber = rand(10000, 99999);
                } while (Student::where('student_number', $studentNumber)->exists());

                $student->name = $request->input('name');
                $student->user_id = $user->id;
                $student->student_number = $studentNumber;
                $student->place_of_birth = $request->input('place_of_birth');
                $student->date_of_birth = $request->input('date_of_birth');
                $student->address = $request->input('address');
                $student->no_handphone = $request->input('no_handphone');
                $student->profile_pic = $request->input('profile_pic');
                $student->created_at = now();
                $student->save();

                return response()->json($successResponse, 200);
            }
        }
    }

    public function createGuru(Request $request)
    {
        // Get the current login siswa
        // $user = Auth::User();

        // for testing only
        $user = \App\Models\User::find(6);

        if($user) {
            $validator = Validator::make($request->all(), [
                'name' => 'required',
                'date_of_birth' => 'required',
                'subjects' => 'required',
                'address' => 'required',
                'no_handphone' => 'required',
                'email' => 'required',
                'employee_number' => 'required',
            ], [
                'name.required' => 'Your name is required',
                'subjects.required' => 'Subjects is required',
                'date_of_birth.required' => 'Your date of birth is required',
                'address.required' => 'Your address is required',
                'no_handphone.required' => 'Your phone number is required',
                'email.required' => 'Your email is required',
                'employee_number.required' => 'Your employee number is required',
            ]);

            $errorResponse = [
                "error" => $validator->errors(),
            ];

            $successResponse = [
                "message" => "Your profile was created"
            ];

            if ($validator->fails()) {
                return response()->json([$errorResponse], 400);
            } else {

                // if user already has a profile
                if (Teacher::where('user_id', $user->id)->exists()) {
                    return response()->json(['error' => 'User already has a teacher profile'], 400);
                }

                $teacher = new Teacher();

                $teacher->name = $request->input('name');
                $teacher->user_id = $user->id;
                $teacher->date_of_birth = $request->input('date_of_birth');
                $teacher->address = $request->input('address');
                $teacher->no_handphone = $request->input('no_handphone');
                $teacher->email = $request->input('email');
                $teacher->employee_number = $request->input('employee_number');
                $teacher->created_at = now();
                $teacher->save();

                return response()->json($successResponse, 200);
            }
        }
    }
}

<?php

namespace App\Http\Requests;

class ClassSubjectAssignmentRequest extends FormRequest
{
    /**
     * Determine if the user is authorized to make this request.
     */
    public function authorize(): bool
    {
        return true;
    }

    /**
     * Get the validation rules that apply to the request.
     *
     * @return array<string, \Illuminate\Contracts\Validation\ValidationRule|array<mixed>|string>
     */
    public function rules(): array
    {
        return [
            'class_subject_id' => 'required|numeric|gt:0',
            'description' => 'required|string',
            'submission_content_type' => 'required',
        ];
    }

    /**
     * @return array
     * Custom validation message
     */
    public function messages(): array
    {
        return [
            'class_subject_id.required' => 'The class subject ID field is required.',
            'class_subject_id.numeric' => 'The class subject ID must be a number.',
            'class_subject_id.gt' => 'The class subject ID must be a positive number.',
            'description.required' => 'The description field is required.',
            'description.string' => 'The description must be a string.',
            'submission_content_type.required' => 'The submission content type field is required.',
        ];
    }
}
